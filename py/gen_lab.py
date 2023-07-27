import argparse
import yaml
from gen.gen_behaviors import *
from gen.gen_routes import *
from gen.gen_policies import *

TOPOPATH="test/kaz"
TOPONAME = "bgp"
RELPATH="kaz"

def get_devices_links():
    devices = {}
    links = {}
    with open(f"{TOPOPATH}/node-list","r") as f:
        nodes = f.readlines()
        for node in nodes:
            n = node.split("#")
            devices[n[0]] = n[2].strip()
    with open(f"{TOPOPATH}/link-info","r") as f:
        nodes = f.readlines()
        for node in nodes:
            n = node.strip().split("#")
            try:
                links[n[0]].append((n[2],n[1],n[3]))
            except:
                links[n[0]]=[(n[2],n[1],n[3])]
    # inits = gen_init(devices,links)
    return devices,links

def get_types():
    with open(f"{TOPOPATH}/link-list","r") as f:
        neighbors = {}
        for link in f.readlines():
            infos = link.strip().split("#")
            as1 = int(infos[0])
            as2 = int(infos[1])
            rel = infos[3]
            r1 = ""
            r2 = ""
            if rel =="P2P" or rel =="NWONKNU" or rel =="UNKNOWN":
                r1 = "peer"
                r2 = "peer"
            elif rel == "P2C":
                r1 = "provider"
                r2 = "customer"
            else:
                r1 = "customer"
                r2 = "provider"
            try:
                neighbors[as1][r2].append(as2)
            except:
                try:
                    neighbors[as1][r2] = [as2]
                except:
                    neighbors[as1] = {}
                    neighbors[as1][r2] = [as2]
            try:
                neighbors[as2][r1].append(as1)
            except:
                try:
                    neighbors[as2][r1] = [as1]
                except:
                    neighbors[as2] = {}
                    neighbors[as2][r1] = [as1]
    return neighbors   

def output(labdata:dict, filename:str):
    with open(filename,"w") as f:
        yaml.dump(labdata,f)

def add_inits(scene: dict,inits):
    if "inits" not in scene.keys():
        scene["inits"] = inits
    else:
        scene["inits"].extend(inits)

def add_policies(policies_conf: dict,policies: list):
    if "policy_deployments" not in policies_conf.keys():
        policies_conf["policy_deployments"] = policies
    else:
        policies_conf["policy_deployments"].extend(policies)

def add_routes(routes_conf: dict,routes: list):
    if "routes" not in routes_conf.keys():
        routes_conf["routes"] = routes
    else:
        routes_conf["routes"].extend(routes)

def add_behaviors(scene: dict,behaviors):
    if "behaviors" not in scene.keys():
        scene["behaviors"] = behaviors
    else:
        scene["behaviors"].extend(behaviors)

# if __name__=="__main__":
#     parser = argparse.ArgumentParser()
#     parser.add_argument("--init",action="store_true")
#     args=parser.parse_args()
#     if args.init:
#         print("hello",end="")
#     pass
scene = {"topo_name":TOPONAME}
reset_scene = {"topo_name":TOPONAME}
devices,links = get_devices_links()
add_inits(scene,gen_beh_init(devices,links))
add_behaviors(reset_scene,gen_beh_reset_all(devices))
routes = {"topo_name":TOPONAME}
policies = {"topo_name":TOPONAME}
add_routes(routes,gen_routes_each_as(devices))
add_policies(policies,gen_commercial_policies(links))
add_policies(policies,gen_reject_invalid_rpki_policies(devices,[]))
scene["routes_path"] = f"{RELPATH}/routes.yaml"
scene["policies_path"] = f"{RELPATH}/policies.yaml"
output(scene,f"{TOPOPATH}/scene.yaml")
output(reset_scene,f"{TOPOPATH}/reset_scene.yaml")
output(routes,f"{TOPOPATH}/routes.yaml")
output(policies,f"{TOPOPATH}/policies.yaml")

# inits = gen_inits_from()
# output({"topo_name":"bgp","inits":inits},f"{TESTPATH}scene.yaml")

# gen_routes_from("bgp")

# for i in range(200,1200,200):
#     temproutes = {"topo_name":TOPONAME}
#     add_routes(temproutes,gen_random_routes(devices,i))
#     output(temproutes,f"{TOPOPATH}/routes_{i}.yaml")

# devices = {}
# with open(f"{TESTPATH}node-list","r") as f:
#     nodes = f.readlines()
#     for node in nodes:
#         n = node.split("#")
#         devices[n[0]] = n[2].strip().split("/")
# behaviors = gen_temp_policies_add(devices)
# output({"topo_name":"bgp","behaviors":behaviors},f"{TESTPATH}/scene_applypolicy.yaml")
