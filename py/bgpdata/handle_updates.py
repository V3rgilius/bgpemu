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
                if "table" not in file:
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
                # ("prefix-any","104.244.42.0/24")
            ])
            # with open(f"{path}/{router}.txt","w") as f:
            while rec := s.get_next_record():
                while elem := rec.get_next_elem():
                    try:
                        aspath = elem.fields['as-path']
                    except:
                        aspath = ""
                    try:
                        nexthop = elem.fields['next-hop']
                    except:
                        nexthop = ""
                    # f.write(f"{rec.time}#{elem.peer_asn}#{router[1:]}#{elem.fields['prefix']}#{elem.fields['as-path']}#{elem.fields['next-hop']}\n")
                    merged_elems.append((rec.time,str(elem.peer_asn),router[1:],elem.fields['prefix'],aspath,nexthop))
                    # merged_elems.append((rec.time,elem.peer_asn,router[1:],elem.fields['prefix'],elem.fields['as-path'],elem.fields['next-hop']))
    # merged_files = [f"{path}/{router}.txt" for router in files]
    random.shuffle(merged_elems)
    return sorted(merged_elems,key=cmp_t)


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
def parse_path(aspath:str):
    return aspath.split(" ")

def get_affected_detail(details):
    t = 0
    result = {}
    for detail in details:
        if t!= detail[0]:
            t = detail[0]
            result[t]=set([detail[1]])
        else:
            result[t].add(detail[1])
    return result


def get_affected_as(linksinfo,elems,prefix:str,attacker) -> list:
    # not_affected = set([str(n) for n in range(1,101)])
    affected = []
    affected_details = []
    for elem in elems:
        if(elem[4]) == "":
            aspath = [elem[1]]
            if elem[1] == attacker:
                continue
        else:
            aspath = parse_path(elem[4])
        if elem[3] == prefix and aspath[-1] == attacker:
            affected.append(elem[1])
            affected_details.append((int(elem[0]),elem[1]))
            if len(linksinfo[elem[2]]) == 1:
                affected.append(elem[2])
                affected_details.append((int(elem[0]),elem[2]))
    return affected,get_affected_detail(affected_details)

def get_path_change(linksinfo,elems,asn,prefix):
    def parse(path):
        return [asn]+path.split(" ")
    paths = []
    result = {}
    flag = False
    if len(linksinfo[asn]) == 1:
        flag=True
    for elem in elems:
        if elem[3] == prefix:
            if (flag and elem[2]==asn) or (elem[1]==asn):
                if elem[4] != "":
                    paths.append((int(elem[0]),elem[4]))
                else:
                    if flag:
                        paths.append((int(elem[0]),elem[1]))
                    else:
                        continue
    temp = get_affected_detail(paths)
    for k in temp.keys():
        result[k] = list(map(parse,temp[k]))
    return paths,result

def get_routing_convergence_times(elems, t):
    convergence_times = []
    start_time = int(elems[0][0])
    end_time = int(elems[0][0])

    for elem in elems:
        if int(elem[0]) - end_time > t:
            convergence_times.append((start_time, end_time))
            start_time = int(elem[0])
        end_time = int(elem[0])
    convergence_times.append((start_time, end_time))
    return convergence_times

def get_packets_count(elems, t, asn):
    packet_counts = []
    start = elems[0][0]
    cnt = 0
    end = start + t

    for elem in elems:
        if elem[1] != asn:
            continue
        while elem[0] > end:
            packet_counts.append((start, end, cnt))
            start = end
            end += t
            cnt = 0
        cnt += 1
    packet_counts.append((start, end, cnt))
    return packet_counts

# elems = merge_each_router("mrts")
# times = get_routing_convergence_times(elems,3)
# for time in times:
#     print(time[1]-time[0])

# print(len(affected))
# with open("mrts/all.txt","w") as f:
#     for elem in elems:
#         f.write(f"{'#'.join(list(map(str,elem)))}")
#         f.write("\n")
    # f.write("\n".join(not_aff))
