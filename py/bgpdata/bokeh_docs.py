# 任意前缀的扩散范围、从某一自治系统到另一自治系统的路径变化、某一路由器发送update报文的情况
from bokeh.plotting import figure
from bokeh.models import Select,HoverTool,LabelSet
from bokeh.layouts import column
from bokeh.application import Application
from bokeh.application.handlers.function import FunctionHandler
from bokeh.server.server import Server
from bokeh.models import ColumnDataSource, Slider
import networkx as nx
from .handle_updates import merge_each_router,get_affected_as
from .show import show_links_raw

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

analyzer = DataAnalyzer("test/real")

def get_spread_doc(doc):
    affected,not_affected = get_affected_as(analyzer.links,analyzer.elems,[])
    p = show_links_raw(list(analyzer.devices.keys()),analyzer.links_list,affected)

    # def callback(attr, old, new):
    #     if new == 0:
    #         data = df
    #     else:
    #         data = df.rolling(f"{new}D").mean()
    #     source.data = ColumnDataSource.from_df(data)

    # slider = Slider(start=0, end=30, value=0, step=1, title="Smoothing by N Days")
    # slider.on_change('value', callback)

    doc.add_root(column(p))

def get_path_change_doc(doc):
    # 创建文档1的内容
    p1 = figure(width=400, height=400, title="Document 1")
    p1.circle([1, 2, 3, 4], [5, 7, 2, 4])
    doc_1 = column(p1)
    doc.add_root(doc_1)



def get_convergence_doc(doc):
    pass

def get_updates_doc(doc):
    pass

