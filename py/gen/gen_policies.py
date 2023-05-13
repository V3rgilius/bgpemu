import argparse
import yaml

TOPOPATH="test/topo117/"

def gen_reject_invalid_rpki_policies(deviceinfos:dict,apply_devices:list) ->dict:
    policies = []
    if apply_devices == []:
        apply_devices = deviceinfos.keys()
    for device in apply_devices:
        policies.append({
            "router_name":f"r{device}",
            "policies": [{
                "name":"rejectrpki",
                "statements":[{
                    "name":"rej",
                    "conditions":{
                        "rpki_result":3
                    },
                    "actions":{
                        'route_action':"REJECT"
                    }
                }]
            }],
            "assignments": [{
                "name":"global",
                "direction":"IMPORT",
                "policies":[{"name":"refuserpki"}],
                "default_action": "REJECT"
            }]
        })
    return policies

# def gen_temp_policies_add(deviceinfos:dict) -> dict:
#     inits = []
#     for d in deviceinfos.keys():
#         tempb = {
#             "name": f"AddRoute r{d}",
#             "device_name": f'r{d}',
#             "steps":[{"name": "addroute","cmds":{
#                 "container": f'r{d}',
#                     "cmds": [
#                         f"gobgp global policy import add refuserpki"
#                     ]}}]
#         }
#         inits.append(tempb)
#     return inits

 
def gen_from_links(neighbors):
    for node in neighbors:
        temp_pd = {
            
        }


def gen_init():
    pass

# gen_from_links()
# if __name__=="__main__":
#     parser = argparse.ArgumentParser()
#     parser.add_argument("--init",action="store_true")
#     args=parser.parse_args()
#     if args.init:
#         print("hello",end="")
#     pass
