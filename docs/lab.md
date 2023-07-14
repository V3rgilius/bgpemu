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
`bgpemu lab deploy <场景文件>` 在已创建的仿真网络中部署实验场景，由场景文件定义，格式如下


## 路由信息
路由信息可通过`bgpemu lab routes deploy <路由信息文件>` 添加到路由器，也可通过场景文件的引用进行添加。路由信息的文件格式如下

## 路由策略
路由策略可通过`bgpemu lab policies deploy <路由策略文件>` 添加到路由器，也可通过场景文件的引用进行添加。路由策略的文件格式如下
