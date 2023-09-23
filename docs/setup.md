# 安装

本节介绍如何配置创建仿真拓扑所需的环境，包含以下两部分
1. 配置依赖
2. 配置K8s集群

# 配置依赖
- 安装Golang
- 安装Docker
- 安装kubeadm

# 配置K8s集群
## 使用kubeadm创建集群
### 1. 配置cri-dockerd
k8s默认使用containerd，需要切换到[cri-dockerd](https://github.com/Mirantis/cri-dockerd/releases/tag/v0.3.1)，下载deb包安装，启用服务。
```
systemctl daemon-reload
systemctl enable cri-docker --now
systemctl is-active cri-docker
```

### 2. 编写config文件
1. 生成初始化配置模板
```
kubeadm token generate
kubeadm config print init-defaults >> init.conf
```
2. 将生成的token填入到配置文件中
    - 修改criSocket为unix:///var/run/cri-dockerd.sock
    - 修改advertiseAddress, token, kubernetesVersion, name等信息
    - 指定podCIDR（使用---与原配置分隔），决定了集群所能运行的最大Pod数量
    ```YAML
    apiVersion: kubeadm.k8s.io/v1beta3
    bootstrapTokens:
    - groups:
        - system:bootstrappers:kubeadm:default-node-token
        token: <Generated token>
        ttl: 24h0m0s
        usages:
        - signing
        - authentication
    kind: InitConfiguration
    localAPIEndpoint:
        advertiseAddress: <Master IP>
        bindPort: 6443
    nodeRegistration:
        criSocket: unix:///var/run/cri-dockerd.sock
        imagePullPolicy: IfNotPresent
        name: master
        taints: null
    ---
    apiServer:
        timeoutForControlPlane: 1m0s
    apiVersion: kubeadm.k8s.io/v1beta3
    certificatesDir: /etc/kubernetes/pki
    clusterName: kubernetes
    controllerManager: {}
    dns: {}
    etcd:
        local:
            dataDir: /var/lib/etcd
    imageRepository: registry.k8s.io
    kind: ClusterConfiguration
    kubernetesVersion: 1.26.3
    networking:
        dnsDomain: cluster.local
        serviceSubnet: 10.96.0.0/12
        podSubnet: 10.244.0.0/16
    scheduler: {}
    ```

### 3. 初始化集群
1. 执行`kubeadm init --config <config_file>`
2. 启用kubectl 
    ```
    mkdir -p $HOME/.kube
    sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
    sudo chown $(id -u):$(id -g) $HOME/.kube/config
    ```

### 4. 添加其他机器作为K8s工作节点
配置过程与上一节基本相同。运行`kubeadm config print join-defaults `生成配置，修改相关条目，然后在worker上执行 `kubeadm join --config <config_file>`

## 使用kubectl配置集群
### 配置[flannel](https://github.com/flannel-io/flannel) cni
`kubectl apply -f https://github.com/flannel-io/flannel/releases/latest/download/kube-flannel.yml`

## 使用kne配置集群
`kne deploy kne/deploy/kne/external-multinode.yaml`
> manifests/metallb/metallb-native.yaml 中的ValidatingWebhookConfiguration全部注释，可能有版本差异问题，导致验证不通过

> deploy/kne/external-multinode.yaml 中meshnet使用grpc可能在跨主机node间无法连接，可选择vxlan

# Trouble Shooting
##  重置集群
- 在master 节点上
清空节点 `kubectl drain --ignore-daemonsets <节点名称>`  
删除节点 `kubectl delete nodes `
- 在每个节点上
    ```
    sudo kubeadm reset --cri-socket unix:///var/run/cri-dockerd.sock
    sudo rm -r /etc/cni/net.d
    rm -r $HOME/.kube/config
    ```
## 重置后出现IP占用问题
问题：cni0" already has an IP address different from xxx
使用kubeadm搭建集群前在每个K8s节点删除cni0和flannel.1设备
```
sudo rm -rf /etc/cni/net.d
sudo rm -rf /var/lib/cni/
sudo ip link set cni0 down
sudo ip link set flannel.1 down
sudo ip link delete cni0
sudo ip link delete flannel.1
```
## 访问grpc服务
从master节点上访问，访问服务ip的50051端口即可

## 修改单个K8s节点上最大Pod数量
1. 打开`/var/lib/kubelet/config.yaml`，在结尾添加`maxPods: 500`，运行`sudo systemctl restart kubelet`重启kubelet服务。
2. 修改init配置，在ClusterConfiguration段中设置所需的node-cidr-mask-size大小，如这里可以设置为15（最多512个Pod）。
[Kubeadm customize](https://kubernetes.io/zh-cn/docs/setup/production-environment/tools/kubeadm/control-plane-flags/#controllermanager-flags)

## 让master运行pod
```bash
kubectl describe node|grep -E "Name:|Taints:"
kubectl taint node master node-role.kubernetes.io/control-plane-
kubectl taint node master node-role.kubernetes.io/control-plane=:NoSchedule
```

