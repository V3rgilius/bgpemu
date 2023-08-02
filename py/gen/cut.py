TOPOPATH="test/kaz"
NEWPATH = "test/kaz2"

def load():
    nodes = {}
    with open(f"{TOPOPATH}/node-list","r") as f:
        for line in f.readlines():
            node = line.strip().split("#")
            nodes[node[0]] = line
    links = []
    with open(f"{TOPOPATH}/link-list","r") as f:
        for line in f.readlines():
            links.append(line.strip().split('#'))
    return nodes,links

def output(nodes,links):
    with open(f"{NEWPATH}/node-list","w") as f:
        for node in nodes:
            f.write(nodes[node])
    with open(f"{NEWPATH}/link-list","w") as f:
        for link in links:
            f.write(link)
nodes,links = load()
newlinks = []
removed = ["58106","701","6939","35104","5511","1299","1239","50509","209141","43727","3216","60299"]
for r in removed:
    del nodes[r]
for link in links:
    if link[0] in removed or link[1] in removed:
        continue
    else:
        newlinks.append("#".join(link)+"\n")
output(nodes,newlinks)