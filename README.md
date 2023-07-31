# bgpemu

<!-- ![banner]()

![badge]()
![badge]()
[![license](https://img.shields.io/github/license/:user/:repo.svg)](LICENSE) -->
[![standard-readme compliant](https://img.shields.io/badge/readme%20style-standard-brightgreen.svg?style=flat-square)](https://github.com/RichardLitt/standard-readme)

bgpemu是一个BGP网络仿真工具

# 目录
- [安装](#安装)
- [使用](#使用)
- [Thanks](#thanks)
- [License](#license)

# 安装
见[文档](docs/setup.md)
# 编译
安装Golang
```
git clone git@github.com:V3rgilius/bgpemu.git
make build 
```

# 使用
bgpemu的使用包括以下几个部分
- [仿真网络](docs/topo.md)
- [场景还原](docs/lab.md)
- [数据收集](docs/data.md)

## 示例
在完成集群的配置，并且编译完成后，可以按以下步骤搭建[示例](examples/kaz)中的仿真网络，并部署场景。
### 创建拓扑
```
bgpemu topo create examples/kaz/kaz.yaml
```
进入网络中的某一容器的shell
```
kubectl exec -it -n bgp r4134 -- sh
```
### 启动BGP路由器
```
bgpemu lab deploy examples/kaz/scene.yaml
```
`scene.yaml`中的配置需按实际情况修改，如rpki的配置
### 添加商业关系路由策略
```
bgpemu lab policies deploy examples/kaz/policies.yaml
```
### 监控BGP报文接收情况
```
bgpemu data start bgp
```
### 启动BGP攻击
```
bgpemu lab routes deploy examples/kaz/atk_routes.yaml
```
### 获取BGP报文数据
报文的转储间隔是30秒，应当至少在这一间隔时间后执行如下命令
```
bgpemu data dump bgp 
```
获得的数据在当前目录的`mrts/`下
### 重置实验
```
bgpemu lab deploy examples/kaz/reset_scene.yaml
```
如果未获取数据可改为`examples/kaz/reset_scene_nodata.yaml`

### 重置后进行RPKI防护实验
```bash
bgpemu lab deploy examples/kaz/scene.yaml
bgpemu lab policies deploy examples/kaz/policies.yaml
bgpemu data start bgp
bgpemu lab deploy examples/kaz/atk_scene_rpki.yaml
# bgpemu data start bgp 后等待30s或更长
bgpemu data dump bgp 
```

# Thanks

本项目主要基于[KNE](https://github.com/openconfig/kne)和[meshnet-cni](https://github.com/networkop/meshnet-cni)的工作

# License