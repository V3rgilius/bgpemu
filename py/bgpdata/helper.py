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

def get_edges_from_paths(paths):
    edges = set()
    for path in paths:
        last = ""
        for node in path:
            if last!="":
                edges.add((last,node))
                edges.add((node,last))
            last = node
    return list(edges)