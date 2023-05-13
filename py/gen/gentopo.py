import random
import networkx as nx
import matplotlib.pyplot as plt
import matplotlib.text as plttext
from matplotlib.backends.backend_svg import FigureCanvasSVG

TOPOPATH="test/topo100/"

import random

def complete_graph(n):
    edges = []
    nodes = list(range(1,n+1))
    for i in nodes:
        for j in range(i + 1, n+1):
            edges.append((i, j))
    return nodes,edges


def generate_random_graph(n, d,i):
    # 创建 n 个节点
    nodes = list(range(1,n+1))
    # 生成一棵树
    tree_edges = generate_tree(nodes,d)
    # 随机添加边，使得每个节点的度数在 [1, d] 之间
    ixs = random.sample(nodes,i)
    for node in ixs:
        degree = random.randint(1, d//2)
        targets = random.sample(nodes, degree)
        if(node in targets):
            continue
        edges = [(node, target) for target in targets if (node, target) not in tree_edges and (target, node) not in tree_edges]
        tree_edges.extend(edges)
    # 返回节点和边列表
    return nodes, tree_edges

def generate_tree(nodes:list,d:int):
    # 生成一棵树
    nodes_copy = nodes[:]
    root = random.choice(nodes)
    tree = []
    queue = [root]
    nodes_copy.remove(root)
    while queue:
        node = queue.pop(0)
        if len(nodes_copy)<d :
            children=nodes_copy[:]
        else:
            children = random.sample(nodes_copy,random.randint(d//2,d))
        if children:
            for child in children:
                tree.append((node, child))
                queue.append(child)
                nodes_copy.remove(child)
    return tree

def count_degree(nodes, edges):
    # 计算每个节点的度数
    degree = {node: 0 for node in nodes}
    for edge in edges:
        degree[edge[0]] += 1
        degree[edge[1]] += 1
    return degree


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
        subnet_mask = "255.255.240.0"
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

# t= generate_network(1000)
# nodes = []
# links = []
# ips=[]
# for node in t:
#     ip = generate_subnet()
#     if(ip in ips):
#         print(f"warning:{ip}")
#     ips.append(ip)
#     nodes.append(f"{node+1}#AS{node+1}#{ip}\n")
#     if node ==0:
#         continue
#     links.append(f"{t[node]['parent']+1}#{node+1}#100M#P2C\n")

# with open(TOPOPATH+"node-list","w") as f:
#     f.writelines(nodes)
# with open(TOPOPATH+"link-list","w") as f:
#     f.writelines(links)

def is_valid(nodes,links):
    linkset = len(set(links))
    nodescopy = nodes[:]
    for link in links:
        if (link[1],link[0]) in links:
            print(link[1],link[0])
        if link[1] in nodescopy:
            nodescopy.remove(link[1])
        else:
            print(link[1])
    print(linkset)

def gen_graph_view(links):
    G = nx.Graph()
    # 添加边到图中
    G.add_edges_from(links)

    # 使用 Force-directed Layout 算法布置节点位置
    pos = nx.kamada_kawai_layout(G)
    fig_size = max(6, len(G.nodes) / 5)
    fig = plt.figure(figsize=(fig_size, fig_size))
    # 绘制节点和边
    labels = {node: f"AS{node}" for node in G.nodes}
    # nx.draw_networkx(G,pos, with_labels=True, labels=labels)
    nx.draw_networkx_nodes(G, pos, node_color='lightblue', node_size=50)
    nx.draw_networkx_edges(G, pos, edge_color='gray', width=0.5,node_size = 1000)
    nx.draw_networkx_labels(G, pos, labels, font_size=10)
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

nodes,links= generate_random_graph(100,10,50)
# nodes,links = complete_graph(50)
nodes_out = []
links_out = []
ips=[] 
gen_graph_view(links)
# is_valid(nodes,links)
for node in nodes:
    ip = generate_subnet()
    nodes_out.append(f"{node}#AS{node}#{ip}\n")
for link in links:
    links_out.append(f"{link[0]}#{link[1]}#100M#P2P\n")
with open(TOPOPATH+"node-list","w") as f:
    f.writelines(nodes_out)
with open(TOPOPATH+"link-list","w") as f:
    f.writelines(links_out)

# with open(TOPOPATH+"link-list","r") as f:
#     lines = f.readlines()
#     links = [line.strip().split("#") for line in lines]
# graph_links = [(int(link[0]),int(link[1])) for link in links]
# gen_graph_view(graph_links)