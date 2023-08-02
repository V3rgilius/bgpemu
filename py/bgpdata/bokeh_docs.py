# 任意前缀的扩散范围、从某一自治系统到另一自治系统的路径变化、某一路由器发送update报文的情况
import datetime
from bokeh.plotting import figure
from bokeh.models import Select,HoverTool,LabelSet
from bokeh.layouts import column,row
from bokeh.application import Application
from bokeh.application.handlers.function import FunctionHandler
from bokeh.server.server import Server
from bokeh.models import ColumnDataSource, Slider
import networkx as nx
# from bokeh.server.util import url_decode
from bgpdata.helper import get_edges_from_paths
from .handle_updates import get_packets_count, merge_each_router,get_affected_as,get_path_change
from .show import show_affected_nodes,show_paths

class DataAnalyzer():
    def __init__(self,topo_path,mrt_path = "mrts"):
        self.devices,self.links,self.links_list = self.get_devices_links(topo_path)
        self.elems = merge_each_router(mrt_path)

    def get_devices_links(self,path):
        devices = {}
        links = {}
        links_list=[]
        with open(f"{path}/link-list", "r") as f:
            lines = f.readlines()
            links_list = [tuple(line.split("#")[:2]) for line in lines]
        with open(f"{path}/node-list","r") as f:
            nodes = f.readlines()
            for node in nodes:
                n = node.split("#")
                devices[n[0]] = n[2].strip()
        with open(f"{path}/link-info","r") as f:
            nodes = f.readlines()
            for node in nodes:
                n = node.strip().split("#")
                try:
                    links[n[0]].append((n[2],n[1],n[3]))
                except:
                    links[n[0]]=[(n[2],n[1],n[3])]
        return devices,links,links_list

analyzer = DataAnalyzer("test/kaz2")

def get_topo_doc(doc):
    p,node_render = show_affected_nodes(list(analyzer.devices.keys()),analyzer.links_list,['7018','3320','3257','6830','3356','2914','5511','3491','1239','6453','6762','1299','12956','701','6461'])
    p.grid.visible = False
    doc.add_root(p)

def get_spread_doc(doc):
    args = doc.session_context.request.arguments
    # "104.244.42.0/24","136168"
    prefix_param = args.get("p")[0].decode()
    asn_param = args.get("asn")[0].decode()
    affected,detail = get_affected_as(analyzer.links,analyzer.elems,prefix_param,asn_param)
    # affected.extend(["23673","23764"])
    timeline = list(detail.keys())
    if(len(timeline)==1):
        timeline.append(timeline[0]+1)
    if(len(timeline)==0):
        timeline = [0,1]
    p,node_render = show_affected_nodes(list(analyzer.devices.keys()),analyzer.links_list,affected)
    p.grid.visible = False
    slider = Slider(title="Timeline", start=timeline[0], end=timeline[-1], step=1, value=timeline[0])
    def callback(attr, old, new):
        # 获取滑块的值
        time = slider.value
        try:
            diffs = detail[time]
        except:
            return
        indexes = node_render.data_source.data["index"]
        fill_colors = ["white" for i in indexes]
        for diff in diffs:
            fill_colors[indexes.index(diff)] = "red"
        node_render.data_source.data["fill_colors"] =fill_colors
    slider.on_change('value', callback)
    doc.add_root(row(p,slider))

def get_path_change_doc(doc):
    args = doc.session_context.request.arguments
    # "104.244.42.0/24","6939"
    prefix_param = args.get("p")[0].decode()
    asn_param = args.get("asn")[0].decode()
    paths,detail = get_path_change(analyzer.links,analyzer.elems,asn_param,prefix_param)
    timeline = list(detail.keys())
    if(len(timeline)==1):
        timeline.append(timeline[0]+1)
    p, line_render = show_paths(list(analyzer.devices.keys()),analyzer.links_list,detail[timeline[0]])
    p.grid.visible = False
    slider = Slider(title="Timeline", start=timeline[0], end=timeline[-1], step=1, value=timeline[0])
    def callback(attr, old, new):
        # 获取滑块的值
        time = slider.value
        try:
            diffs = get_edges_from_paths([detail[time][0]])
        except:
            return
        indexes = line_render.data_source.data["index"]
        fill_colors = ["black" for i in indexes]
        for diff in diffs:
            if diff in indexes:
                fill_colors[indexes.index(diff)] = "red"
            elif diff[::-1] in indexes:
                fill_colors[indexes.index(diff[::-1])] = "red"
        line_render.data_source.data["colors"] =fill_colors
    slider.on_change('value', callback)
    doc.add_root(row(p,slider))

def get_convergence_doc(doc):
    pass

def get_updates_doc(doc):
    # asn = "44356"
    args = doc.session_context.request.arguments
    asn_param = args.get("asn")[0].decode()
    cnts = get_packets_count(analyzer.elems,1,asn_param)
    timestamps = []
    ys = []
    for cnt in cnts:
        timestamps.append((cnt[0]+cnt[1])/2)
        ys.append(cnt[2])
    xs = [datetime.datetime.utcfromtimestamp(ts).strftime('%H:%M:%S') for ts in timestamps] 
    # 创建数据源
    source = ColumnDataSource(data=dict(x=xs, y=ys))

    # 创建柱状图
    p = figure(height=500, x_range=xs)
    vbar_render = p.vbar(x='x', top='y', width=0.5, source=source)
    labels = LabelSet(x='x', y='y', text='y', level='glyph',
                    x_offset=-10, y_offset=0, source=source)
    p.add_layout(labels)
    # 创建滑块
    slider = Slider(title='Time Range', start=1, end=10, value=1, step=1)

    def update_data(attr, old, new):
        # 获取滑块的值
        sample_length = slider.value
        cnts = get_packets_count(analyzer.elems,sample_length,asn_param)
        timestamps = []
        ys = []
        for cnt in cnts:
            timestamps.append((cnt[0]+cnt[1])/2)
            ys.append(cnt[2])
        # 根据滑块值更新数据源
            xs = [datetime.datetime.utcfromtimestamp(ts).strftime('%H:%M:%S') for ts in timestamps] 
        updated_data = dict(x=xs, y=ys)
        vbar_render.data_source.data = updated_data
        vbar_render.glyph.width = sample_length/2

    # 监听滑块值的变化
    slider.on_change('value', update_data)
    # 创建布局
    layout = column(slider, p)
    doc.add_root(layout)

