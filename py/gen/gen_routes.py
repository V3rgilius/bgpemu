import random
def generate_subnet():
    while True:
        # 随机生成四个字节的 IP 地址
        ip = ".".join(str(random.randint(0, 255)) for _ in range(4))
        # 检查 IP 地址是否是私有地址
        if is_private_ip(ip):
            continue
        # 随机生成子网掩码长度（24 到 30 位之间）
        subnet_length = 24
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

def gen_random_routes(deviceinfos:dict,n:int):
    routes = {}
    devices = list(deviceinfos.keys())
    for i in range(n):
        prefix = generate_subnet()
        device = random.choice(devices)
        try:
            routes[f"r{device}"].append({
                    "nlri":{
                        "prefix_len":prefix.split("/")[1],
                        "prefix":prefix.split("/")[0]
                    },
                })
        except:
            routes[f"r{device}"]=[{
                    "nlri":{
                        "prefix_len":prefix.split("/")[1],
                        "prefix":prefix.split("/")[0]
                    },
                }]
    
    result = []
    for device in routes:
        result.append({
            "name":device,
            "paths":routes[device]
        })
    return result
        


def gen_routes_each_as(deviceinfos:dict) ->dict :
    routes = []
    for device in deviceinfos:
        routes.append({
            "name":f"r{device}",
            "paths":[
                {
                    "nlri":{
                        "prefix_len":deviceinfos[device].split("/")[1],
                        "prefix":deviceinfos[device].split("/")[0]
                    },
                }
            ]
        })
    return routes

def gen_routes_selected(deviceinfos:dict,selected:list) ->dict :
    routes = []
    for device in selected:
        routes.append({
            "name":f"r{device}",
            "paths":[
                {
                    "nlri":{
                        "prefix_len":deviceinfos[device].split("/")[1],
                        "prefix":deviceinfos[device].split("/")[0]
                    },
                }
            ]
        })
    return routes

# def gen_temp_routes_add(deviceinfos:dict) -> dict:
#     inits = []
#     for d in deviceinfos.keys():
#         tempb = {
#             "name": f"AddRoute r{d}",
#             "device_name": f'r{d}',
#             "steps":[{"name": "addroute","cmds":{
#                 "container": f'r{d}',
#                     "cmds": [
#                         f"gobgp global rib add -a ipv4 {'/'.join(deviceinfos[d])} origin igp"
#                     ]}}]
#         }
#         inits.append(tempb)
#     return inits


# def gen_routes_from(topo_name):
#     devices = {}
#     with open(f"{TESTPATH}node-list","r") as f:
#         nodes = f.readlines()
#         for node in nodes:
#             n = node.split("#")
#             devices[n[0]] = n[2].strip().split("/")
#     routes = gen_routes(devices)
#     routed = {
#         "topoName":topo_name,
#         # "behaviors":gen_temp_routes_add(devices)
#         # "routes":routes,
#         "policy_deployments":gen_policies(devices)
#     }
#     output(routed,f"{TESTPATH[:-1]}routes.yaml")