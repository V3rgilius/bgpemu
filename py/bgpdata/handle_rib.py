from _pybgpstream import BGPStream, BGPRecord, BGPElem


# 创建BGPStream对象
stream = BGPStream()

# 设置MRT文件路径
stream.set_data_interface("singlefile")
stream.set_data_interface_option("singlefile", 'upd-file', "test/testmrt/r2/20230506.1551.updates.dump")
# 启动BGPStream
stream.start()

# 遍历BGP消息
while rec := stream.get_next_record():
    print(f'{rec.time}')
    while elem := rec.get_next_elem():
        pass
        print(f"{elem.fields['prefix']}\t{elem.fields['as-path']}\t{elem.fields['next-hop']}")

