```
$ bgpemu lab
Lab commands.

Usage:
  bgpemu lab [command]

Available Commands:
  deploy      Deploy a lab
  generate    Generate a lab scene from a topo file
  policies    Policy commands.
  routes      Route commands.

Flags:
  -h, --help   help for lab
```
## 实验场景
`bgpemu lab deploy <场景文件>` 在已创建的仿真网络中部署实验场景，由场景文件定义，格式如下，部分字段为[GoBGP API](https://github.com/osrg/gobgp/blob/master/api/gobgp.proto)定义的类型。
字段|类型|含义
--- | --- | ---
topo_name | string | 拓扑名称
inits | []Behavior | 初始化行为列表
routes_path | string | 路由信息文件路径
policies_path | string | 路由策略文件路径
behaviors | []Behavior | 行为列表

**Behavior**
字段|类型|含义
--- | --- | ---
name | string | 行为名称，用于日志输出
device_name | string | 设备Pod名称
is_async | bool | 是否异步执行
steps | []Step | 步骤列表

**Step**
字段|类型|含义
--- | --- | ---
name | string | 步骤名称
body |以下类型的一种| 步骤
- Commands cmds 
- Wait wait 
- FileTrans file 
- StartBgpStep sbs 
- AddPeerStep aps 

**Commands**
字段|类型|含义
--- | --- | ---
container | string | 执行命令的容器名称
cmds | []string | 命令列表

**Wait**
字段|类型|含义
--- | --- | ---
time | uint32 | 等待一段时间，单位是毫秒
timestamp | uint64 | 等待至时间戳

**FileTrans**
字段|类型|含义
--- | --- | ---
src | string | 源文件/目录路径(与kubectl cp 所用格式一致)
des | string | 目的文件/目录路径

**StartBgpStep**
字段|类型|含义
--- | --- | ---
global | gobgp.Global | GoBGP的全局设置
rpki | gobgp.AddRpkiRequest | RPKI设置

**AddPeerStep**
字段|类型|含义
--- | --- | ---
peers | []gobgp.Peer | Peer列表

## 路由信息
路由信息可通过`bgpemu lab routes deploy <路由信息文件>` 添加到路由器，也可通过场景文件的引用进行添加。路由信息的文件格式如下，部分字段为[GoBGP API](https://github.com/osrg/gobgp/blob/master/api/gobgp.proto)定义的类型。
字段|类型|含义
--- | --- | ---
topo_name | string | 拓扑名称
routes | []Route | 路由配置列表

Route
字段|类型|含义
--- | --- | ---
name | string | 设备名称
mrt_path | string | 
paths | []BgpPath | BGP路由信息列表


BgpPath
字段|类型|含义
--- | --- | ---
nlri | gobgp.IPAddressPrefix | IPv4前缀
origin | OriginType | 来源类型
aspath | []gobgp.AsSegment | As段列表
next_hop | string | 下一跳IPv4地址
local_pref | uint32 | 本地优先级
med | uint32 | MED
communities | []uint32 | 社区列表
    


## 路由策略
路由策略可通过`bgpemu lab policies deploy <路由策略文件>` 添加到路由器，也可通过场景文件的引用进行添加。路由策略的文件格式如下，部分字段为[GoBGP API](https://github.com/osrg/gobgp/blob/master/api/gobgp.proto)定义的类型。

PolicyDeployments
字段|类型|含义
--- | --- | ---
topo_name | string | 拓扑名称
policy_deployments | []PolicyDeployment | 路由策略列表

PolicyDeployment
字段|类型|含义
--- | --- | ---
router_name | string | 设备名称
defined_sets | []gobgp.DefinedSet | 
statements | []gobgp.Statement | 
policies | []gobgp.Policy | 
assignments | []gobgp.PolicyAssignment | 
peer_groups | []gobgp.PeerGroup | 


