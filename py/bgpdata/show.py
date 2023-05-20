from bokeh.models import HoverTool,LabelSet
from bokeh.plotting import figure
import networkx as nx
from bokeh.models import ColumnDataSource
from bokeh.plotting import figure

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

def show_links_raw(nodes,links,diffs):
    # 创建一个简单的有向图
    G = nx.DiGraph()
    G.add_nodes_from(nodes)
    G.add_edges_from(links)
    # 创建一个 Bokeh 图形对象
    plot = figure(title="Interactive Network", x_range=(-2.1, 2.1), y_range=(-2.1, 2.1),width=1000, height=1000,)
    plot.axis.visible = False
    # plot.toolbar.logo = None
    # plot.toolbar_location = None
    # 提取节点位置
    pos = nx.shell_layout(G)
    # pos = nx.drawing.nx_agraph.pygraphviz_layout(G,prog="dot")
    # 绘制节点
    # x,y = zip(*pos.values())
    indexes = list(pos.keys())
    x = [pos[i][0] for i in indexes]
    y = [pos[i][1] for i in indexes]
    neighbors = get_neighbors(links)
    neighbor_source = [",".join(neighbors[i]) for i in indexes]
    fill_colors = ["white" for i in indexes]
    for diff in diffs:
        fill_colors[indexes.index(diff)] = "red"
    source = ColumnDataSource(data=dict(
        x=x,
        y=y,
        index=indexes,
        neighbors=neighbor_source,
        fill_colors = fill_colors
    ))
    node_source = plot.circle('x','y',size=20, fill_color='fill_colors', line_width=2,source = source)
    # colors = node_source.glyph.fill_colors
    # 绘制边
    edge_xs, edge_ys = [], []
    for u, v in G.edges():
        x0, y0 = pos[u]
        x1, y1 = pos[v]
        edge_xs.append([x0, x1])
        edge_ys.append([y0, y1])
    plot.multi_line(edge_xs, edge_ys, line_alpha=0.8, line_width=1)

    # 添加交互式工具和提示
    hover = HoverTool(renderers=[node_source], tooltips=[("Node", "@index"),("Neighbors","@neighbors")])
    plot.add_tools(hover)

    labels = LabelSet(x='x', y='y',text='index', level='glyph', source=source,x_offset=5,y_offset=5, )
    plot.add_layout(labels)
    return plot
