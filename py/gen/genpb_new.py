from collections import deque
import random
import yaml
import toml


TOPOPATH = "test/topo60/"
OUTPATH = "test/"
TOPONAME = "topo60"
K8SNODES = 3


def load_node():
    lines = []
    with open(TOPOPATH+"node-list", "r") as f:
        lines = f.readlines()
    nodes = [line.strip().split("#") for line in lines]

    def mapping(a, b):
        return (a, b)
    ret = dict(map(mapping, [node[0]
               for node in nodes], [node[2] for node in nodes]))
    return ret


def load_link():
    lines = []
    with open(TOPOPATH+"link-list", "r") as f:
        lines = f.readlines()
    links = [line.split("#") for line in lines]
    return links


def gen_my_node(asn, ip, globalip, ethlinks,relations):
    links: dict = ethlinks[asn]
    eths = links.keys()
    ips ={}
    for eth in eths:
        ips[eth] = f"{globalip[f'{asn}:{eth}']}/30"
    shells = ["/usr/local/bin/gobgpd > /dev/null 2> /dev/null &"]
    frrshells = [ ]
    nodedata = {"name": f"r{asn}",
                "type": "BGPNODE",
                "ip_addr":ips,
                "config": {
                        "tasks": [{"container": f"r{asn}-frr", "cmds": frrshells},{"container": f"r{asn}", "cmds": shells}],
                        "extra_images": {
                            f"r{asn}-frr": "frrouting/frr:v8.1.0",
                        },
                        "share_volumes": [
                            "zebra"
                        ],
                        "container_volumes":{
                            f"r{asn}-frr":{"volumes":["zebra"],"paths":["/var/run/frr"]},
                            f"r{asn}":{"volumes":["zebra"],"paths":["/var/run/frr"]}
                        }
                },
                "services":
                    {50051: {
                        "name": "gobgp",
                        "inside": 50051
                    }}
                }
    neighbors=[f"{asn}#{links[eth].split(':')[0]}#{globalip[links[eth]]}#{relations[asn][links[eth].split(':')[0]]}\n" for eth in eths]
    with open(f"{TOPOPATH}link-info","a") as f:
        f.writelines(neighbors)
    return nodedata


def gen_node(asn, ip, globalip, ethlinks,relations):
    links: dict = ethlinks[asn]
    eths = links.keys()
    shells = [
        f"ip addr add {globalip[f'{asn}:{eth}']}/30 dev {eth}" for eth in eths]
    shells.append("/usr/local/bin/gobgpd > /dev/null 2> /dev/null &")  # -f /config/gobgp.toml
    # shells.append("sleep 0.5")
    # shells.append(f"gobgp global rib add -a ipv4 {nodes[asn]} origin igp")
    # frrshells = [ "/usr/lib/frr/frrinit.sh start"]
    frrshells = [ ]
    nodedata = {"name": f"r{asn}",
                "vendor": "GOBGP",
                "config": {
                        # "config_path": "/",
                        # "config_file": "gobgp.conf",
                        # "file": f"r{asn}.toml",
                        "tasks": [{"container": f"r{asn}-frr", "cmds": frrshells},{"container": f"r{asn}", "cmds": shells}],
                        "extra_images": {
                            f"r{asn}-frr": "frrouting/frr:v8.1.0",
                        },
                        "share_volumes": [
                            "zebra"
                        ],
                        "container_volumes":{
                            f"r{asn}-frr":{"volumes":["zebra"],"paths":["/var/run/frr"]},
                            f"r{asn}":{"volumes":["zebra"],"paths":["/var/run/frr"]}
                        }
                        # "affinity": affinities[asn]
                },
                # "interfaces":{}
                "services":
                    {50051: {
                        "name": "gobgp",
                        "inside": 50051
                    }}
                }
    # config = {
    #     "global": {
    #         "config": {
    #             "as": int(asn),
    #             "router-id": ip.split("/")[0],
    #         }
    #     },
    #     "neighbors":
    #         [{"config": {"neighbor-address": globalip[links[eth]],
    #                      "peer-as": links[eth].split(":")[0]}} for eth in eths],
    #     "zebra": {
    #         "config": {
    #             "enabled": True,
    #             # "url": "tcp:localhost:2600",
    #             "url": "unix:/var/run/frr/zserv.api",
    #             "redistribute-route-type-list": ["connect"],
    #             "version": 6
    #         }
    #     }
    # }
    # with open(f"{OUTPATH}r{asn}.toml", 'w') as f:
    #     toml.dump(config, f)

    # with open("output/addip.sh",'a') as f:
    #     f.writelines(shells)
    #     f.write(f"kubectl exec -it r{asn} -n bgp -- gobgp global rib add -a ipv4 {nodes[asn]} origin egp\n")
    neighbors=[f"{asn}#{links[eth].split(':')[0]}#{globalip[links[eth]]}#{relations[asn][links[eth].split(':')[0]]}\n" for eth in eths]
    with open(f"{TOPOPATH}link-info","a") as f:
        f.writelines(neighbors)
    return nodedata


# def get_neighbors(asn, nodes, links):
#     neighbors = []
#     for link in links:
#         if (link[0] == asn):
#             neighbors.append((link[1], nodes[link[1]]))
#         if (link[1] == asn):
#             neighbors.append((link[0], nodes[link[0]]))
#     return neighbors


def gen_link(a_asn, z_asn, cnt0, cnt1):
    linkdata = {
        "a_node": f"r{a_asn}",
        "a_int": f"eth{cnt0}",
        "z_node": f"r{z_asn}",
        "z_int": f"eth{cnt1}",
    }
    return linkdata

def gen_ip(ipcnt:int,ip:str)->str:
    bs=ip.split(".")
    bs[3]=str(ipcnt%256)
    bs[2]=str(ipcnt//256+int(bs[2]))
    return ".".join(bs)

def get_relation(relations:dict,link):
    if link[0] not in relations.keys():
        relations[link[0]]={}
    if link[1] not in relations.keys():
        relations[link[1]]={}
    relations[link[0]][link[1]]=link[2].strip()
    relations[link[1]][link[0]]=link[2].strip()[::-1]

def get_neighbors(links):
    neighbors = {}
    for link in links:
        try:
            neighbors[link[0]].append(link[1])
        except:
            neighbors[link[0]]=[link[1]]
        try:
            neighbors[link[1]].append(link[0])
        except:
            neighbors[link[1]]=[link[0]]
    return neighbors

def get_affinities(nodes,links):
    nodes_set = set(nodes)
    neighbors = get_neighbors(links)
    k8s_num = [len(nodes)//K8SNODES for i in range(K8SNODES)]
    k8s_num[-1]+= len(nodes)%K8SNODES
    init_k8s_num = k8s_num[:]
    tags = {}
    flag =True
    while flag:
        idx = k8s_num.index(max(k8s_num))
        cnt = 0
        node_list = [random.choice(list(nodes_set))]
        tag = str(random.randint(1,10000))
        while node_list!= []:
            node = node_list.pop(0)
            tags[node] = {"t":tag}
            nodes_set.remove(node)
            cnt+=1
            for neighbor in neighbors[node]:
                if neighbor not in node_list and neighbor in nodes_set:
                    node_list.append(neighbor)
            if(cnt >= k8s_num[idx]):
                node_list = []
        k8s_num[idx] -= cnt
        cnt = 0
        for i in range(K8SNODES):
            if(init_k8s_num[i] == k8s_num[i]):
                flag = True
                break
            else:
                flag = False
    for node in list(nodes_set):
        for n in neighbors[node]:
            if n in tags.keys():
                tags[node] = tags[n]
                break
    return tags


def gen_topo(nodes, acts, links,kne = False):
    eth_cnt = {}
    global_ip = {}
    eth_links = {}
    ipcnt = {}
    relations = {}
    affinities = get_affinities(acts,links)
    with open(f"{TOPOPATH}link-info","w") as f:
         pass
    for n in acts:
        eth_links[n] = {}
        eth_cnt[n] = 0
        ipcnt[n]=1
    topodata = {"name": "bgp", "nodes": [], "links": []}
    for link in links:
        eth_cnt[link[0]] += 1
        eth_cnt[link[1]] += 1
        ethcnt0 = eth_cnt[link[0]]
        ethcnt1 = eth_cnt[link[1]]
        subnetsize0 = int(nodes[link[0]].split("/")[1])
        subnetsize1 = int(nodes[link[1]].split("/")[1])
        if(subnetsize0>=subnetsize1):
            global_ip[f"{link[0]}:eth{ethcnt0}"] = gen_ip(ipcnt[link[0]],nodes[link[0]].split("/")[0])
            ipcnt[link[0]] += 1
            global_ip[f"{link[1]}:eth{ethcnt1}"] = gen_ip(ipcnt[link[0]],nodes[link[0]].split("/")[0])
            ipcnt[link[0]] += 3
        else:
            global_ip[f"{link[0]}:eth{ethcnt0}"] = gen_ip(ipcnt[link[1]],nodes[link[1]].split("/")[0])
            ipcnt[link[1]] += 1
            global_ip[f"{link[1]}:eth{ethcnt1}"] = gen_ip(ipcnt[link[1]],nodes[link[1]].split("/")[0])
            ipcnt[link[1]] += 3

        eth_links[link[0]][f"eth{ethcnt0}"] = f"{link[1]}:eth{ethcnt1}"
        eth_links[link[1]][f"eth{ethcnt1}"] = f"{link[0]}:eth{ethcnt0}"
        topodata["links"].append(gen_link(link[0], link[1], ethcnt0, ethcnt1))
        get_relation(relations,link)

    for node in acts:
        if kne:
            topodata["nodes"].append(
                gen_node(node, nodes[node], global_ip, eth_links,relations))
        else:
            topodata["nodes"].append(
                gen_my_node(node, nodes[node], global_ip, eth_links,relations))
        eth_cnt[node] = 0
    if kne:
        with open(f"{TOPOPATH}{TOPONAME}_kne.yaml", 'w') as f:
            yaml.dump(topodata, f)
    else:
        with open(f"{TOPOPATH}{TOPONAME}.yaml", 'w') as f:
            yaml.dump(topodata, f)


if __name__ == "__main__":
    # with open("{OUTPATH}addip.sh", 'w') as f:
    #     pass
    nodes = load_node()
    links = load_link()
    a_nodes = [link[0] for link in links]
    z_nodes = [link[1] for link in links]
    act_nodes = set(a_nodes).union(set(z_nodes))
    gen_topo(nodes, act_nodes, links)
    gen_topo(nodes, act_nodes, links,kne=True)
    # affinities = get_affinities(nodes,links)
    # with open("testdata","w") as f:
    #     for a in affinities:
    #         f.write(str(affinities[a])+"\n")
