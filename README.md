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

# Thanks

本项目主要基于[KNE](https://github.com/openconfig/kne)和[meshnet-cni](https://github.com/networkop/meshnet-cni)的工作

# License