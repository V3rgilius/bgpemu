

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