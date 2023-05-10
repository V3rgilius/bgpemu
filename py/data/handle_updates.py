from _pybgpstream import BGPStream, BGPRecord, BGPElem
import os 
import random

def get_files(path):
    files = {}
    routers = os.listdir(path)
    for router in routers:
        if "." not in router:
            files[router]=[]
            for file in os.listdir(f"{path}/{router}"):
                files[router].append(f"{path}/{router}/{file}")
    return files

# def merge_all(files: 'list[str]'):
#     # (timestamp,send_asn,recv_asn,prefix,aspath,nexthop)
#     file_handles = [open(f,"r") for f in files]
#     cur_elems = [() for f in files]
#     cur_time = 0
#     finished = 0
#     num = len(files)
#     finished_flags = dict.fromkeys(files,False)

#     while finished < num:
              
        

def merge_each_router(path):
    def cmp_t(x):
        return x[0]
    files = get_files(path)
    merged_elems = []
    for router in files:
        for f in files[router]:
            s = get_stream(f,[
            ])
            with open(f"{path}/{router}.txt","w") as f:
                while rec := s.get_next_record():
                    while elem := rec.get_next_elem():
                        f.write(f"{rec.time}#{elem.peer_asn}#{router[1:]}#{elem.fields['prefix']}#{elem.fields['as-path']}#{elem.fields['next-hop']}\n")
                        merged_elems.append((rec.time,elem.peer_asn,router[1:],elem.fields['prefix'],elem.fields['as-path'],elem.fields['next-hop']))
    merged_files = [f"{path}/{router}.txt" for router in files]
    random.shuffle(merged_elems)
    return merged_files,sorted(merged_elems,key=cmp_t)


def get_stream(path,filters:list):
    # 创建BGPStream对象
    stream = BGPStream()
    # 设置MRT文件路径
    stream.set_data_interface("singlefile")
    stream.set_data_interface_option("singlefile", 'upd-file', path) # "test/testmrt/20230506.1551.updates.dump"
    for filter in filters:
        stream.add_filter(*filter)
    # 启动BGPStream
    stream.start()
    # 遍历BGP消息
    return stream

def get_affected_as(elems,prefixes:'list[str]') -> list:
    pass

files, elems = merge_each_router("test/testmrt")
get_affected_as(elems,[
    
])