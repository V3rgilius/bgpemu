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
export_ints|map[string]InternalInterface>|外部接口名称
update_topo|string|更新拓扑所基于的拓扑，仅当更新时使用

**Node**
字段|类型|含义
--- | --- | ---
name | string
type | Type
path | string
ip_addr | map[string]string
services | map[uint32]Service
config | Config

**Link**
字段|类型|含义
--- | --- | ---
a_node| string | 
a_int |string |
z_node| string |
z_int |string |

**InternalInterface**
字段|类型|含义
--- | --- | ---
node | string
node_int | string

**Config**
字段|类型|含义
--- | --- | ---
tasks | []Task
extra_images | map<string, string>
share_volumes | []string
container_volumes | map<string,PublicVolumes>
image | string
affinity | map<string,string>

**PublicVolumes**
字段|类型|含义
--- | --- | ---
volumes | []string
paths | []string

**Task**
字段|类型|含义
--- | --- | ---
container | string
cmds | []string

### 子拓扑

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
