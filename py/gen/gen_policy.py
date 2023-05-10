import argparse
import yaml

TESTPATH="test/topo117/"


def load_links(filename):
    neighbors = {}
    with open(filename,"r") as f:
        for link in f.readlines():
            infos = link.strip().split("#")
            as1 = int(infos[0])
            as2 = int(infos[1])
            rel = infos[3]
            r1 = ""
            r2 = ""
            if rel =="P2P":
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
 
def gen_from_links():
    neighbors = load_links(f"{TESTPATH}link-list")
    for node in neighbors:
        temp_pd = {
            
        }


def gen_init():
    pass

gen_from_links()
# if __name__=="__main__":
#     parser = argparse.ArgumentParser()
#     parser.add_argument("--init",action="store_true")
#     args=parser.parse_args()
#     if args.init:
#         print("hello",end="")
#     pass