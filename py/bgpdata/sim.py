# 令 totalBGPNum＝100
# 设定30％的RPKI部署率，即30％的IP前缀受到RPKI的保护，30％的BGP启用RPKI扩展，此时拓扑图设为G
# 令i＝0／／i的值，代表使用slurm的BGP路由器数量
# 计算图G中，每个点到受害者的最短距离，得到数组D［totalBGPNum］
# 声明res［totalBGPNum］
# fori =0 to (totalBGPNum *0.3):
#     从图G中去掉使用slurm的bgp，得到子图G＇
#     声明D0［totalBGPNum］，初始化为INT＿MAX
#     计算图G＇中，每个点到攻击者的最短距离，更新列表D0［totalBGPNum］
#     forj＝0 to totalBGPNum：／／统计有多少点，举例攻击者更近
#         if DO[j]<D[j]:
#             res[i]++ 
# 返回res 
import show
from collections import deque
import random
from bokeh.io import save


TOPOPATH="test/real"
RPKINUM = 80
TESTNUM =50

def dijkstra_shortest_paths(graph, nodes, start):
    distances = {node: float('inf') for node in nodes}  # 存储每个节点到起点的最短距离
    distances[start] = 0  # 起点到自身的距离为0
    paths = {start: []}  # 存储每个节点到起点的最短路径
    visited = set()  # 存储已访问过的节点

    while len(visited) < len(nodes):
        min_distance = float('inf')
        min_node = None

        # 选择距离最小的节点作为当前节点
        for node in nodes:
            if node not in visited and distances[node] < min_distance:
                min_distance = distances[node]
                min_node = node

        visited.add(min_node)

        # 更新当前节点的邻居节点的距离和路径
        for edge in graph:
            if min_node in edge:
                neighbor = edge[1] if edge[0] == min_node else edge[0]
                new_distance = distances[min_node] + 1
                if new_distance < distances[neighbor]:
                    distances[neighbor] = new_distance
                    paths[neighbor] = paths[min_node] + [min_node]

    return distances, paths


def load_link():
    lines = []
    with open(TOPOPATH+"/link-list", "r") as f:
        lines = f.readlines()
    links = [line.split("#") for line in lines]
    links = [(link[0],link[1]) for link in links]
    return links

def remove_rpki(links,rpkis,attacker):
    new_links = []
    for link in links:
        if link[0] in rpkis or link[1] in rpkis:
            continue
        else:
            new_links.append(link)
    return get_largest_connected_graph(new_links,attacker) 

# def dfs(graph, visited, node):
#     visited.add(node)
#     for edge in graph:
#         if node in edge and edge[1] not in visited:
#             neighbor = edge[1] if edge[0] == node else edge[0]
#             dfs(graph, visited, neighbor)

# def get_largest_connected_graph(graph, n):
#     visited = set()
#     dfs(graph, visited, n)
#     largest_graph = [edge for edge in graph if edge[0] in visited and edge[1] in visited]
#     return visited,largest_graph

def get_largest_connected_graph(graph, n):
    visited = set()
    queue = deque([n])  # 使用队列进行广度优先搜索
    largest_graph = []  # 存储最大连通图的边

    while queue:
        node = queue.popleft()
        if node not in visited:
            visited.add(node)
            for edge in graph:
                if node in edge:
                    neighbor = edge[1] if edge[0] == node else edge[0]
                    if neighbor not in visited:
                        queue.append(neighbor)
                        largest_graph.append(edge)

    return visited,largest_graph


def get_affected(nodes,links,rpkis,attacked,attacker):
    d1,p1 = dijkstra_shortest_paths(links,nodes,attacked)
    removed_nodes,removed_links = remove_rpki(links,rpkis,attacker)
    d2,p2 = dijkstra_shortest_paths(removed_links,removed_nodes,attacker)
    affecteds = []
    for node in d2:
        if d1[node]>d2[node]:
            affecteds.append(node)
    return affecteds,p2

def test(attacked,attacker):
    result = []
    links = load_link()
    nodes = set()
    # attacked = "55"
    # attacker = "29"
    rpkis = [attacked]
    for link in links:
        nodes.add(link[0])
        nodes.add(link[1])
    for n in range(RPKINUM):
        # rpkis.append(random.choice(list(nodes.difference(set(rpkis)).difference(set([attacker])))))
        rpkis = [attacked]+random.sample(list(nodes.difference(set([attacker,attacked]))),n)
        affected,paths = get_affected(nodes,links,rpkis,attacked,attacker)
        # show.show_links_raw(list(nodes),links,affected)
        # print(len(affected))
        result.append(len(affected))
    return result



attacked = "13414"
attacker = "136168"
rpkis =[attacked]
links = load_link()
nodes = set()
for link in links:
    nodes.add(link[0])
    nodes.add(link[1])
affected,paths = get_affected(nodes,links,rpkis,attacked,attacker)
plot = show.show_links_raw(list(nodes),links,affected)
save(plot,'network.html')



# rand_att_test = []
# fix_att_test = []
# fixed_attacker = "55"
# fixed_attacked = "29"
# for i in range(TESTNUM):
#     attacker,attacked = random.sample(list(range(1,100)),2)
#     print("Rand: ",attacker,attacked)
#     rand_att_test.append((attacker,attacked,test(str(attacked),str(attacker))))
#     fix_att_test.append(test(fixed_attacked,fixed_attacker))

# # fix_results = []
# rnd_results = []
# for j in range(RPKINUM):
#     # fix_results.append(0)
#     rnd_results.append(0)
#     for i in range(TESTNUM):
#         rnd_results[j]+=rand_att_test[i][2][j]
#         # fix_results[j]+=fix_att_test[i][j]
#     rnd_results[j]/=TESTNUM
#     # fix_results[j]/=TESTNUM


# # print(fix_results)
# with open("testdata","w") as f:
#     f.write("\n".join(list(map(str,rnd_results))))

