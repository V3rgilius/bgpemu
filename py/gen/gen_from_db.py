import random
import sqlite3
import networkx as nx
import matplotlib.pyplot as plt
import matplotlib.text as plttext
from matplotlib.backends.backend_svg import FigureCanvasSVG

tier1List=['7018','3320','3257','6830','3356','2914','5511','3491','1239','6453','6762','1299','12956','701','6461']
searched= set(tier1List)
myas = ['136168','13414','132132','61292','18106','23673']
TOPOPATH="topo/sample/"

def db_connect():
    db = sqlite3.connect('linksdata.db')
    return db

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

def rev_search(src,dsts):
    db = db_connect()
    cursor = db.cursor()
    searched = {src,}
    queue = [src]
    ret_paths =[]
    while queue!=[]:
        node = queue.pop(0)
        neighbors,tmplinks = get_providers(node,cursor)
        for neighbor in neighbors:
            if neighbor in dsts:
                queue=[]
                break
            if neighbor not in searched:
                queue.append(neighbor)
                searched.add(neighbor)
        ret_paths.extend(tmplinks)
    cursor.close()
    db.close()
    return ret_paths 

def get_top(paths,dsts):
    def get_one(t):
        return t[0]
    def get_two(t):
        return t[1]
    providers = list(map(get_one,paths))
    customers = list(map(get_two,paths))
    ret = set()
    for p in providers:
        if p not in customers and p not in dsts:
            ret.add(p)
    return ret

def clean(paths,dsts):
    nouse = get_top(paths,dsts)
    cleaned = []
    for path in paths:
        if path[0] not in nouse:
            cleaned.append(path)
    return cleaned
    
def rev_gen(srcs,dsts):
    paths = {}
    for src in srcs:
        paths[src] = rev_search(src,dsts)
        for top in get_top(paths[src],dsts):
            more = rev_search(top,dsts)
            cnt = len(more)
            more =clean(more,dsts)
            while(len(more)!= cnt):
                cnt = len(more)
                more = clean(more,dsts)
            paths[src].extend(clean(more,dsts))
    return paths


def gen_graph_view(str_links):
    def s2n(t):
        return (int(t[0]),int(t[1]))
    links = list(map(s2n,str_links))
    G = nx.Graph()
    # 添加边到图中
    G.add_edges_from(links)

    # 使用 Force-directed Layout 算法布置节点位置
    # pos = nx.spring_layout(G)
    pos = nx.drawing.nx_agraph.graphviz_layout(G,prog="dot")
    fig_size = max(6, len(G.nodes) / 5)
    fig = plt.figure(figsize=(fig_size, fig_size))
    # 绘制节点和边
    labels = {node: f"AS{node}" for node in G.nodes}
    # nx.draw_networkx(G,pos, with_labels=True, labels=labels)
    nx.draw_networkx_nodes(G, pos, node_color='lightblue', node_size=50)
    nx.draw_networkx_edges(G, pos, edge_color='gray', width=0.5,node_size=1000)
    nx.draw_networkx_labels(G, pos, labels, font_size=4)
    # 显示图像
    elems = fig.findobj()
    for elem in elems:
        if isinstance(elem, plttext.Text):
            elem.set_gid(elem.get_text())
    plt.axis('off')
    canvas = FigureCanvasSVG(fig)
    # 将图形保存为 SVG 文件
    canvas.print_svg(f"{TOPOPATH[:-1]}.svg")
    # with open("mygraph.svg", "w") as f:
    #     f.write(svg_output)
    # plt.savefig('my_graph.png')

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


def output(paths:'dict[str,tuple[str,str]]'):
    nodes= set()
    links= set()
    for src in paths:
        tmp=set(paths[src])
        links=links.union(tmp)
    with open("topo/sample/link-list","w") as f:
        for link in links:
            nodes.add(link[0])
            nodes.add(link[1])
            f.write(f"{link[0]}#{link[1]}#P2C\n")
    sql_1 = "select prefix from origin where ASN=?;"
    db = db_connect()
    cursor = db.cursor()
    with open("topo/sample/node-list","w") as f:
        for node in nodes:
            if node in tier1List:
                print(node)
            cursor.execute(sql_1,(node,))
            try:
                prefix=cursor.fetchall()[0][0]
                if prefix=="0.0.0.0/0":
                    prefix = generate_subnet()
            except:
                prefix=generate_subnet()
            f.write(f"{node}#nil#{prefix}\n")
    cursor.close()
    db.close()
    gen_graph_view(links)

paths = rev_gen(myas,tier1List)
output(paths)