# deviceinfos[asn]=ip
# linkinfos[asn]=[(peer_ip,peer_asn)]

def gen_beh_reset_all(deviceinfos:dict) -> dict:
    inits = []
    for d in deviceinfos.keys():
        tempb = {
            "name": f"Reset r{d}",
            "device_name": f'r{d}',
            "steps":[{"name": "restart","cmds":{
                "container": f'r{d}',
                    "cmds": [
                        "kill -9 $(pidof gobgpd)",
                        "sleep 0.5",
                        "/usr/local/bin/gobgpd > /dev/null 2> /dev/null &",
                        "rm -r /tmp/log/*"
                    ]}}]
        }
        inits.append(tempb)
    return inits

def gen_beh_init(deviceinfos:dict,linkinfos:dict) -> dict:
    inits = []
    for d in deviceinfos.keys():
        tempb = {
            "name": f"StartRouter{d}Bgp",
            "device_name": f'r{d}',
            "steps":[{"name": "start","sbs":{
                    "global": {
                        "asn": int(d),
                        "router_id": deviceinfos[d].split("/")[0],
                    },
                    "rpki":{
                        "address": "10.10.21.110",
                        "port":3323
                    }
            }}]
        }
        tempb["steps"].append({
            "name": "addpeer",
            "aps":{
                "peers": [{
                    "conf":{
                        "neighbor_address": peer[0],
                        "peer_asn": int(peer[1])
                    }
                } for peer in linkinfos[d]]
            }
        })
        inits.append(tempb)
    return inits