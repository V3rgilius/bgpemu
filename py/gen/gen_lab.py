import argparse
import yaml

TESTPATH="test/topo100/"
# deviceinfos[asn]=ip
# linkinfos[asn]=[(peer_ip,peer_asn)]
def gen_init(deviceinfos:dict,linkinfos:dict) -> dict:
    inits = []
    for d in deviceinfos.keys():
        tempb = {
            "name": f"StartRouter{d}Bgp",
            "device_name": f'r{d}',
            "steps":[{"name": "start","sbs":{
                    "global": {
                        "asn": int(d),
                        "router_id": deviceinfos[d],
            }}}]
        }
        for peer in linkinfos[d]:
            tempb["steps"].append({
                "name": "addpeer",
                "aps":{
                    "peer": {
                        "conf":{
                            "neighbor_address": peer[0],
                            "peer_asn": int(peer[1])
                        }
                    }
                }
            })
        inits.append(tempb)
    return inits


def gen_inits_from():
    devices = {}
    links = {}
    with open(f"{TESTPATH}node-list","r") as f:
        nodes = f.readlines()
        for node in nodes:
            n = node.split("#")
            devices[n[0]] = n[2][:-4]
    with open(f"{TESTPATH}linkinfo","r") as f:
        nodes = f.readlines()
        for node in nodes:
            n = node.strip().split("#")
            try:
                links[n[0]].append((n[2],n[1]))
            except:
                links[n[0]]=[(n[2],n[1])]
    inits = gen_init(devices,links)
    return inits

def gen_routes(deviceinfos:dict) ->dict :
    routes = []
    for device in deviceinfos:
        routes.append({
            "name":f"r{device}",
            "paths":[
                {
                    "nlri":{
                        "prefix_len":deviceinfos[device][1],
                        "prefix":deviceinfos[device][0]
                    },
                }
            ]
        })
    return routes

def gen_temp_routes_add(deviceinfos:dict) -> dict:
    inits = []
    for d in deviceinfos.keys():
        tempb = {
            "name": f"AddRoute r{d}",
            "device_name": f'r{d}',
            "steps":[{"name": "addroute","cmds":{
                "container": f'r{d}',
                    "cmds": [
                        f"gobgp global rib add -a ipv4 {'/'.join(deviceinfos[d])} origin igp"
                    ]}}]
        }
        inits.append(tempb)
    return inits

def gen_routes_from(topo_name):
    devices = {}
    with open(f"{TESTPATH}node-list","r") as f:
        nodes = f.readlines()
        for node in nodes:
            n = node.split("#")
            devices[n[0]] = n[2].strip().split("/")
    routes = gen_routes(devices)
    routed = {
        "topo_name":topo_name,
        # "behaviors":gen_temp_routes_add(devices)
        "routes":routes
    }
    output(routed,f"{TESTPATH[:-1]}.yaml")

def output(labdata:dict, filename:str):
    with open(filename,"w") as f:
        yaml.dump(labdata,f)

# if __name__=="__main__":
#     parser = argparse.ArgumentParser()
#     parser.add_argument("--init",action="store_true")
#     args=parser.parse_args()
#     if args.init:
#         print("hello",end="")
#     pass

inits = gen_inits_from()
output({"topo_name":"bgp","inits":inits},f"{TESTPATH}scene.yaml")

# gen_routes_from("bgp")