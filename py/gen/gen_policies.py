import argparse
import yaml

TOPOPATH = "test/topo117/"


def gen_reject_invalid_rpki_policies(deviceinfos: dict, apply_devices: list) -> dict:
    policies = []
    if apply_devices == []:
        apply_devices = deviceinfos.keys()
    for device in apply_devices:
        policies.append({
            "router_name": f"r{device}",
            "policies": [{
                "name": "rejectrpki",
                "statements": [{
                    "name": "rej",
                    "conditions": {
                        "rpki_result": 3
                    },
                    "actions": {
                        'route_action': "REJECT"
                    }
                }]
            }],
            "assignments": [{
                "name": "global",
                "direction": "IMPORT",
                "policies": [{"name": "refuserpki"}],
                "default_action": "REJECT"
            }]
        })
    return policies


def gen_commercial_policies(linksinfo: dict) -> dict:
    policies = []
    for device in linksinfo:
        prvds = []
        peers = []
        cstms = []
        for peer in linksinfo[device]:
            if (peer[2] == "C2P"):
                prvds.append(peer[0]+"/30")
            elif peer[2] == "P2P":
                peers.append(peer[0]+"/30")
            else:
                cstms.append(peer[0]+"/30")
        def_sets = [{
                "defined_type": "COMMUNITY",
                "name": "peer-prov",
                "list": ["65000:10", "65100:20"]
            }]
        def_sets.append({
            "defined_type": "NEIGHBOR",
            "name": "providers-set",
            "list": prvds
        })
        def_sets.append({
            "defined_type": "NEIGHBOR",
            "name": "peers-set",
            "list": peers
        })
        def_sets.append({
            "defined_type": "NEIGHBOR",
            "name": "customers-set",
            "list": cstms
        })
        policies.append({
            "router_name": f"r{device}",
            "defined_sets": def_sets,
            "statements": [{
                "name": "noexport",
                "conditions": {
                    "community_set": {
                        "type": "ANY",
                        "name": "peer-prov"
                    },
                    "neighbor_set": {
                        "type": "INVERT",
                        "name": "customers-set"
                    }
                },
                "actions": {
                    "route_action": "REJECT"
                }
            },
                {
                "name": "providers-in",
                "conditions": {
                    "neighbor_set": {
                        "type": "ANY",
                        "name": "providers-set"
                    }
                },
                "actions": {
                    "community": {
                        "type": "ADD",
                        "communities": ["65100:10"]
                    },
                    "local_pref":{
                        "value": 10
                    }
                }
            },
            {
                "name": "peers-in",
                "conditions": {
                    "neighbor_set": {
                        "type": "ANY",
                        "name": "peers-set"
                    }
                },
                "actions": {
                    "community": {
                        "type": "ADD",
                        "communities": ["65100:20"]
                    },
                    "local_pref":{
                        "value": 20
                    }
                }
            },
            {
                "name": "customers-in",
                "conditions": {
                    "neighbor_set": {
                        "type": "ANY",
                        "name": "customers-set"
                    }
                },
                "actions": {
                    "community": {
                        "type": "ADD",
                        "communities": ["65100:30"]
                    },
                    "local_pref":{
                        "value": 30
                    }
                }
            }],
            "policies": [{
                "name": "peer-prov-export",
                "statements": [{
                    "name": "noexport",
                }]
            },
            {
                "name": "providers-import",
                "statements": [{
                    "name": "providers-in",
                }]
            },
            {
                "name": "peers-import",
                "statements": [{
                    "name": "peers-in",
                }]
            },
            {
                "name": "customers-import",
                "statements": [{
                    "name": "customers-in",
                }]
            }],
            "assignments": [{
                "name": "global",
                "direction": "IMPORT",
                "policies": [{"name": "providers-import"}, {"name": "peers-import"}, {"name": "customers-import"}],
                "default_action": "ACCEPT"
            },
                {
                "name": "global",
                "direction": "EXPORT",
                "policies": [{"name": "peer-prov-export"}],
                "default_action": "ACCEPT"
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
