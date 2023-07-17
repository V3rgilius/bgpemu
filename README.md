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
在完成集群的配置，并且编译完成后，可以按以下步骤搭建[示例](examples/attack_test)中的仿真网络，并部署场景。
### 创建拓扑
```
bgpemu topo create examples/attack_test/bgptopo.yaml
```
进入网络中的某一容器的shell
```
kubectl exec -it -n bgp r13414 -- sh
```
### 启动BGP路由器
```
bgpemu lab deploy examples/attack_test/scene.yaml
```
### 添加路由策略
```
bgpemu lab policies deploy examples/attack_test/policies.yaml
```
### 获得BGP报文接收情况数据
```
bgpemu data start bgp
```
报文的转储间隔是30秒，当获取足够的数据后
```
bgpemu data dump bgp 
```

# Thanks

本项目主要基于[KNE](https://github.com/openconfig/kne)和[meshnet-cni](https://github.com/networkop/meshnet-cni)的工作

# License