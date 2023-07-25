import random
import sqlite3
import networkx as nx
import matplotlib.pyplot as plt
import matplotlib.text as plttext
from matplotlib.backends.backend_svg import FigureCanvasSVG
from gen_from_db import output
myass = ['8393','4134','58106']
tier1List=['7018','3320','3257','6830','3356','2914','5511','3491','1239','6453','6762','1299','12956','701','6461']
searched= set(tier1List)
myas = ['136168','13414','132132','61292','4844','18106','23673']
TOPOPATH="test/kaz"
MAXHOPS = 3


def db_connect():
    db = sqlite3.connect('py/gen/linksdata.db')
    return db

def get_bgppeers(n,cursor):
    def get_one(t:tuple):
        return (t[0],t[1])
    def get_two(t:tuple):
        return (t[0],t[1][::-1])
    sql_1 = "select as2,type from aslinks_with_type where as1=?;"
    sql_2 = "select as1,type from aslinks_with_type where as2=?;"
    peers = set()
    cursor.execute(sql_1,(n,))
    peers = peers.union(set(map(get_one,cursor.fetchall())))
    cursor.execute(sql_2,(n,))
    peers = peers.union(set(map(get_two,cursor.fetchall())))
    return peers


def get_providers(n,cursor):
    def get_one(t:tuple):
        return t[0]
    sql_1 = "select as2 from aslinks_with_type where as1=? and type=?;"
    sql_2 = "select as1 from aslinks_with_type where as2=? and type=?;"
    neighbors = set()
    cursor.execute(sql_1,(n,"c2p"))
    neighbors = neighbors.union(set(map(get_one,cursor.fetchall())))
    cursor.execute(sql_2,(n,"p2c"))
    neighbors = neighbors.union(set(map(get_one,cursor.fetchall())))
    links = []
    for neighbor in neighbors:
        links.append((neighbor,n))
    return neighbors,links


def get_top(paths,dsts):
    tier1_set = set(dsts)
    ret = set()
    customers = set()
    ret = set()
    for path in paths:
        ret.add(path[0])
        customers.add(path[1])
    ret.difference_update(tier1_set)
    ret.difference_update(customers)
    return ret

def clean(paths,dsts):
    nouse = get_top(paths,dsts)
    cleaned = []
    for path in paths:
        if path[0] not in nouse:
            cleaned.append(path)
    return cleaned

def rev_search_all(src,dsts):
    db = db_connect()
    cursor = db.cursor()
    searched = {src,}
    found = set()
    queue = [(src,0)]
    hops = 0
    ret_paths =[]
    while queue!=[]:
        node = queue.pop(0)
        if queue == [] or node[1]!=hops:
            hops+=1
            if hops > MAXHOPS:
                break
        neighbors,tmplinks = get_providers(node[0],cursor)
        for neighbor in neighbors:
            if neighbor in dsts:
                found.add(neighbor)
            if neighbor not in searched:
                queue.append((neighbor,hops+1))
                searched.add(neighbor)
        ret_paths.extend(tmplinks)
    cur_len = len(ret_paths)+1
    while cur_len!= len(ret_paths):
        cur_len = len(ret_paths)
        ret_paths = clean(ret_paths,dsts)
    cursor.close()
    db.close()
    return ret_paths 

def rev_gen_all(srcs,dsts):
    paths = {}
    for src in srcs:
        paths[src] = rev_search_all(src,dsts)   # 得到到tier1的路径列表
    links= set()
    links_type = set()
    nodes = set()
    db = db_connect()
    cursor = db.cursor()
    for src in paths:
        tmp=set(paths[src])
        links=links.union(tmp)
    for link in links:
        nodes.add(link[0])
        nodes.add(link[1])
        links_type.add((*link,"p2c"))

    for node in nodes:  # 获得的AS拓扑中的节点之间如果有连接则添加到拓扑中
        peers = get_bgppeers(node,cursor)
        for peer in peers:
            if peer[0] in nodes:
                if (peer[0],node) not in links:
                    links.add((node,peer[0]))
                    links_type.add((node,*peer))
    cursor.close()
    db.close()
    return links,links_type

def gen_graph_view(str_links):
    def s2n(t):
        return (int(t[0]),int(t[1]))
    links = list(map(s2n,str_links))
    G = nx.Graph()
    # 添加边到图中
    G.add_edges_from(links)

    pos = nx.drawing.nx_agraph.pygraphviz_layout(G,prog="dot")
    fig_size = max(6, len(G.nodes) / 5)
    fig = plt.figure(figsize=(fig_size, fig_size))
    # 绘制节点和边
    labels = {node: f"AS{node}" for node in G.nodes}
    edge_labels = {link: "P" for link in links}
    nx.draw_networkx_nodes(G, pos, node_color='lightblue', node_size=50)
    nx.draw_networkx_edges(G, pos, edge_color='gray', width=0.5,node_size=200)
    nx.draw_networkx_labels(G, pos, labels, font_size=4)
    elems = fig.findobj()
    for elem in elems:
        if isinstance(elem, plttext.Text):
            elem.set_gid(elem.get_text())
    plt.axis('off')
    canvas = FigureCanvasSVG(fig)
    canvas.print_svg(f"{TOPOPATH}/topo.svg")

def generate_subnet():
    while True:
        # 随机生成四个字节的 IP 地址
        ip = ".".join(str(random.randint(0, 255)) for _ in range(4))
        # 检查 IP 地址是否是私有地址
        if is_private_ip(ip):
            continue
        # 随机生成子网掩码长度（24 到 30 位之间）
        subnet_length = 20
        # 计算子网掩码
        subnet_mask = "255.255.255.0"
        # 计算网络地址
        network_address = ".".join(str(int(ip.split(".")[i]) & int(subnet_mask.split(".")[i]))
                                   for i in range(4))
        # 返回 IP 地址和子网掩码
        return f"{network_address}/{subnet_length}"

def is_private_ip(ip):
    # 检查 IP 地址是否是私有地址
    octets = ip.split(".")
    return (octets[0] == "10" or
            (octets[0] == "172" and 16 <= int(octets[1]) <= 31) or
            (octets[0] == "192" and octets[1] == "168"))

def output(links):
    nodes= set()
    with open(f"{TOPOPATH}/link-list","w") as f:
        for link in links:
            nodes.add(link[0])
            nodes.add(link[1])
            f.write(f"{link[0]}#{link[1]}#{link[2].upper()}\n")
    sql_1 = "select prefix from origin where ASN=?;"
    sql_2 = "select name from asname where asn=?;"
    db = db_connect()
    cursor = db.cursor()
    with open(f"{TOPOPATH}/node-list","w") as f:
        for node in nodes:
            # if node in tier1List:
            #     print(node)
            cursor.execute(sql_1,(node,))
            try:
                prefix=cursor.fetchall()[0][0]
                if ":" in prefix:
                    prefix = "0.0.0.0/0"
            except:
                prefix="0.0.0.0/0"
            cursor.execute(sql_2,(node,))
            name = cursor.fetchall()[0][0]
            f.write(f"{node}#{name}#{prefix}\n")
    cursor.close()
    db.close()

paths,paths_type = rev_gen_all(myass,tier1List)
output(paths_type)

