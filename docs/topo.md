## 创建拓扑

在完成对集群的配置后，bgpemu使用home下的默认kubeconfig对集群进行配置，使用`bgpemu topo create`命令创建拓扑
```
$ bgpemu topo create
Create Topology
Usage:
  bgpemu topo create <拓扑文件>
```
###  拓扑文件格式
拓扑文件使用YAML描述，格式如下
字段|类型|含义
--- | --- | ---
name|string|拓扑名称
nodes|[]Node|网络节点列表
links|[]Link|网络链路列表
export_ints|map[string]InternalInterface>|外部接口名称到内部节点接口的映射关系
update_topo|string|更新拓扑所基于的拓扑，仅当更新时使用

**Node**
字段|类型|含义
--- | --- | ---
name | string | 节点名称
type | Type | 节点类型，值包括BGPNODE、HOST、SUBTOPO
path | string | 节点拓扑文件(当节点类型为SUBTOPO时有效)
ip_addr | map[string]string | 接口名到IP地址的映射关系
services | map[uint32]Service | 端口号到K8s服务的映射关系
config | Config | 节点的其他配置

**Link**
字段|类型|含义
--- | --- | ---
a_node| string | 起始节点名称
a_int |string | 起始节点接口名称
z_node| string | 结束节点名称
z_int |string | 结束节点接口名称

**InternalInterface**
字段|类型|含义
--- | --- | ---
node | string | 内部接口所属节点
node_int | string | 内部接口对应节点上的接口名称

**Config**
字段|类型|含义
--- | --- | ---
tasks | []Task | 任务列表
extra_images | map[string]string | 节点所属Pod上运行的其他容器，容器名到镜像标签的映射
share_volumes | []string | 节点所属Pod上共享卷的名称列表
container_volumes | map[string]PublicVolumes | Pod中容器的共享卷设置，键是容器名
image | string | 指定Pod的镜像标签
affinity | map[string]string | Pod的Pod亲和性设置，详见[PodAffinity](https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node/#inter-pod-affinity-and-anti-affinity)

**PublicVolumes**
字段|类型|含义
--- | --- | ---
volumes | map[string]string | 共享卷在容器上的挂载目录

**Task**
字段|类型|含义
--- | --- | ---
container | string | 执行命令的容器名
cmds | []string | 命令列表

### 子拓扑
TODO
## 删除拓扑
```
$ bgpemu topo delete
Delete Topology
Usage:
  bgpemu topo delete <拓扑文件>
```
删除拓扑时需输入创建所用的拓扑文件路径，若有更新拓扑的操作，输入最后一次添加的拓扑文件，并且拓扑文件中定义了所基于的拓扑。
## 更新拓扑
TODO
```
$ bgpemu topo update
Update Topology
Usage:
  bgpemu topo update <拓扑文件>
```
