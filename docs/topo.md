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
update_topo|string|更新拓扑所基于的拓扑路径，仅当更新时使用

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
is_resilient | bool | 节点是否可以在创建后动态添加端口，以便拓扑的弹性变化
envs | map[string]string | Pod的Envs，自动在Pod中所有容器上配置
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
在Node的类型设置为SUBTOPO，然后引用外部的拓扑文件，外部的拓扑文件应该定义相应的外部端口
## 删除拓扑
```
$ bgpemu topo delete
Delete Topology
Usage:
  bgpemu topo delete <拓扑文件>
```
删除拓扑时需输入创建所用的拓扑文件路径，若有更新拓扑的操作，输入最后一次添加的拓扑文件，并且拓扑文件中定义了所基于的拓扑。
## 更新拓扑
```
$ bgpemu topo update
Update Topology
Usage:
  bgpemu topo update <拓扑文件>
```
用于更新的拓扑文件与普通拓扑文件格式相同，且支持子拓扑。
在Nodes中定义新创建的节点，与创建拓扑时的格式相同。
在Links中定义需要添加的链路，按照节点的类型可以将链路分为三种：
- 链路两端的节点均已在原拓扑上创建
- 链路一端的节点在原拓扑上，另一端是新创建的节点
- 链路两端的节点均是新创建的节点
不论是哪种类型，对于在原拓扑上的节点，都需要在创建时将`config`中的`is_resilient`标志设置为`true`；而对于新创建的节点，则无要求。
> 新创建的节点会使用`ip_addr`中的IP地址分配给端口，而当为原拓扑上的节点新增端口时，新增的端口不会自动分配IP地址，需要手动为端口添加（可以创建一个场景文件完成这件事）。

首次创建的仿真拓扑及其之后若干的更新拓扑形成的是一条链，当为这条链添加新的更新拓扑配置时，`bgpemu`会沿着这条链还原出原有的拓扑，并将新拓扑附加到其上。
在删除时，应该指定最近一次更新的拓扑文件作为输入，`bgpemu`会沿着链还原完整拓扑，并进行与单个拓扑文件情形下相同的删除操作。