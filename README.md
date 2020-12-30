
## 微服务
>微服务架构是一种架构模式，它提倡将单一应用程序划分成一组小的服务，服务之间相互协调、互相配合，为用户提供最终价值。每个服务运行在其独立的进程中，服务和服务之间采用轻量级的通信机制相互沟通（通常是基于HTTP的Restful API).每个服务都围绕着具体的业务进行构建，并且能够被独立的部署到生产环境、类生产环境等。另外，应尽量避免统一的、集中的服务管理机制，对具体的一个服务而言，应根据业务上下文，选择合适的语言、工具对其进行构---- Martin Fowler

+ 优点
  + 提升开发交流，每个服务足够内聚，足够小，代码容易理解
  + 方便进行多语言开发，系统不会被长期限制在某个技术栈上
  + 服务独立测试、部署、升级、发布，扩展性强
  + 容易扩大开发团队，可以针对每个服务（service）组件开发团队
  + 提高容错性（fault isolation），一个服务的内存泄露并不会让整个系统瘫痪
+ 缺点
  + 微服务提高了系统和编排的复杂度
  + 服务之间的分布式通信问题
  + 服务的注册与发现问题
  + 服务之间的分布式事务&&一致性问题

> <a href="https://www.zdnet.com/article/8-ways-to-test-your-readiness-for-microservices/" />[知乎]什么是微服务-华为云的回答</a>

+ 微服务的前置条件
  + 模块划分清晰，功能保持独立
  + 所有交互通过api进行
  + 有特别强的运维团队
  + 良好的自动化测试
  + 完善的监控系统

> <a href="https://www.zdnet.com/article/8-ways-to-test-your-readiness-for-microservices/" />8-ways-to-test-your-readiness-for-microservices</a>

---

## 架构分析
<img src="https://github.com/siyoo/devops/blob/master/images/整体架构.png" height="400px" width="400px" />

1. ETCD

<img src="https://github.com/siyoo/devops/blob/master/images/etcd.jpg" height="400px" width="400px" />

+ HTTP Server： 用于处理用户发送的API请求以及其它etcd节点的同步与心跳信息请求
+ Store：用于处理etcd支持的各类功能的事务，包括数据索引、节点状态变更、监控与反馈、事件处理与执行等等，是etcd对用户提供的大多数API功能的具体实现
+ Raft：Raft强一致性算法的具体实现，是etcd的核心
> <a href="http://thesecretlivesofdata.com/raft/" />Raft算法演示</a>
+ WAL：Write Ahead Log（预写式日志），是etcd的数据存储方式。除了在内存中存有所有数据的状态以及节点的索引以外，etcd就通过WAL进行持久化存储。WAL中，所有的数据提交前都会事先记录日志。Snapshot是为了防止数据过多而进行的状态快照；Entry表示存储的具体日志内容
> <a href="https://blog.csdn.net/bbwangj/article/details/82584988" />[CSDN] ETCD 简介</a>

2. 透明网关
+ 初始化调用链, 生成jaeger Span Context
  + opentracing: OpenTracing 是一套标准，它通过提供平台无关、厂商无关的API，使得开发人员能够方便的添加（或更换）追踪系统的实现，标准中目前定义了两种类型的引用:
    + ChildOf: 可以理解为是继承关系
    + FollowsFrom: 顺序关系，并无继承行为
  <img src="https://github.com/siyoo/devops/blob/master/images/jaeger.png" height="400px" width="800px" />

+ jaeger:受Dapper和OpenZipkin启发的Jaeger是由Uber作为开源发布的分布式跟踪系统。它用于监视和诊断基于微服务的分布式系统
    + jaeger-client：jaeger 的客户端，实现了opentracing协议；
    + jaeger-agent：jaeger client的一个代理程序，client将收集到的调用链数据发给agent，然后由agent发给collector；
    + jaeger-collector：负责接收jaeger client或者jaeger agent上报上来的调用链数据，然后做一些校验，比如时间范围是否合法等，最终会经过内部的处理存储到后端存储；
  + Trace表示对一次请求完整调用链的跟踪(可能不止两个)，两个服务的请求/响应过程叫做一次Span, 一个span一般会记录这个调用单元内部的一些信息，例如：
    + 日志信息
    + 标签信息
    + 开始/结束时间
+ 鉴权
+ 设置meta data
+ 根据ETCD获取后端服务地址，调用后端

## CI/CD
1. Jenkins file大体内容（groovy语言）
  <img src="https://github.com/siyoo/devops/blob/master/images/jenkins.png" height="800px" width="400px" />

2. k8s
   1. Kubernetes，是一种容器编排技术，它可以帮助用户省去应用容器化过程的许多手动部署和扩展操作
   2. k8s支持的容器
      + Docker
      + Rkt
      + Containerd
      + Lxd
      + Kata Container
      + 其他容器技术
   > <a href="https://landscape.cncf.io/" /> CNCF:云原生计算基金会</a>
   3. k8s基础知识
      1. 基本概念
         + Pod 是 Kubernetes 的最小工作单元，每个 Pod 包含一个或多个容器。Pod 中的容器会作为一个整体被 Master 调度到一个 Node 上运行
         + Node是k8s集群中相对于Master而言的工作主机。Node可以是一台物理主机，也可以是一台虚拟机
         + Namespace可以在一个Kubernetes集群中创建多个“虚拟集群”，这些Namespace之间可以完全隔离
      2. k8s架构 (<a href="https://blog.tianfeiyu.com/" /> 田飞雨的博客 </a>)
        <img src="https://github.com/siyoo/devops/blob/master/images/k8s架构.png" height="1000px" width="1000px" />

      + apiserver 是 kubernetes 中与 etcd 直接交互的一个组件，其控制着 kubernetes 中核心资源的变化。它主要提供了以下几个功能：

        提供 Kubernetes API，包括认证授权、数据校验以及集群状态变更等，供客户端及其他组件调用  
        代理集群中的一些附加组件组件，如 Kubernetes UI、metrics-server、npd 等  
        创建 kubernetes 服务，即提供 apiserver 的 Service，kubernetes Service  
        资源在不同版本之间的转换
      + scheduler 的目的就是为每一个 pod 选择一个合适的 node，整体流程可以概括为三步

        获取未调度的 podList  
        通过调度算法为 pod 选择一个合适的 node  
        提交数据到 apiserver

      + Controller Manager作为集群内部的管理控制中心，负责集群内的Node、Pod副本、服务端点（Endpoint）、命名空间（Namespace）、服务账号（ServiceAccount）、资源定额（ResourceQuota）的管理，当某个Node意外宕机时，Controller Manager会及时发现并执行自动化修复流程，确保集群始终处于预期的工作状态。  
        每个Controller通过API Server提供的接口实时监控整个集群的每个资源对象的当前状态，当发生各种故障导致系统状态发生变化时，会尝试将系统状态修复到“期望状态”。

        <img src="https://github.com/siyoo/devops/blob/master/images/cm.png" height="400px" width="400px" />

        + Replication Controller简称RC,称为副本控制器。副本控制器的作用即保证集群中一个RC所关联的Pod副本数始终保持预设值
        + Node Controller通过API Server实时获取Node的相关信息，实现管理和监控集群中的各个Node节点的相关控制功能
        + ResourceQuota Controller 确保指定的资源对象在任何时候都不会超量占用系统物理资源  
          1）容器级别：对CPU和Memory进行限制  
          2）Pod级别：对一个Pod内所有容器的可用资源进行限制  
          3）Namespace级别：包括  
            a. Pod数量  
            b. Replication Controller数量  
            c. Service数量  
            d. ResourceQuota数量  
            e. Secret数量  
            f. 可持有的PV（Persistent Volume）数量
        + Namespace Controller 用户通过API Server可以创建新的Namespace并保存在etcd中，Namespace Controller定时通过API Server读取这些Namespace信息
        + Endpoint Controller 负责生成和维护所有Endpoints对象的控制器。它负责监听Service和对应的Pod副本的变化
          + Endpoints表示了一个Service对应的所有Pod副本的访问地址
        + Service Controller 集群与外部的云平台之间的一个接口控制器。Service Controller监听Service变化，如果是一个LoadBalancer类型的Service，则确保外部的云平台上对该Service对应的LoadBalancer实例被相应地创建、删除及更新路由转发表
        > <a href="https://www.cnblogs.com/Su-per-man/p/10942856.html" />Controller Manager简介</a>
      + kube-proxy: service 实际的路由转发都是由 kube-proxy 组件来实现的，service 仅以一种 VIP（ClusterIP） 的形式存在，kube-proxy 主要实现了集群内部从 pod 到 service 和集群外部从 nodePort 到 service 的访问，kube-proxy 的路由转发规则是通过其后端的代理模块实现的
        + kube-proxy 的代理模块目前有四种实现方案，userspace、iptables、ipvs、kernelspace
      + kubelet 是运行在每个节点上的主要的“节点代理”，每个节点都会启动 kubelet进程，用来处理 Master 节点下发到本节点的任务，按照 PodSpec 描述来管理Pod 和其中的容器（PodSpec 是用来描述一个 pod 的 YAML 或者 JSON 对象)

        kubelet 通过各种机制（主要通过 apiserver ）获取一组 PodSpec 并保证在这些 PodSpec 中描述的容器健康运行  
        kubelet 默认监听四个端口:
        10250（kubelet API）：kubelet server 与 apiserver 通信的端口，定期请求 apiserver 获取自己所应当处理的任务，通过该端口可以访问获取 node 资源以及状态

        10248（健康检查端口）：通过访问该端口可以判断 kubelet 是否正常工作, 通过 kubelet 的启动参数 --healthz-port 和 --healthz-bind-address 来指定监听的地址和端口

        4194（cAdvisor 监听）：kublet 通过该端口可以获取到该节点的环境信息以及 node 上运行的容器状态等内容，访问 http://localhost:4194 可以看到 cAdvisor 的管理界面,通过 kubelet 的启动参数

        10255 （readonly API）：提供了 pod 和 node 的信息，接口以只读形式暴露出去，访问该端口不需要认证和鉴权

   4. deployment file大体内容

       + 一个Ingress - 流量从集群外部流入到集群内部的你的服务上
       + 一个Service -内部负载均衡器，用于将流量路由到内部的Pod上
       + 一个Deployment - Pod 和 ReplicaSet 之上，提供了一个声明式定义（declarative）方法，将 Pod 和 ReplicaSet 的实际状态改变到目标状态


   5. 上面各组件的关系

       + Ingress的 servicePort 应该匹配service的 port  
         Ingress的 serviceName 应该匹配服务的 name

       + 连接Deployment和Service  
         Service selector应至少与Pod的一个标签匹配  
         Service的targetPort应与Pod中容器的containerPort匹配  
         Service的port可以是任何数字。多个Service可以使用同一端口号，因为它们被分配了不同的IP地址

3. Docker
   1. 简介
      + docker 最开始使用linux container（LXC）进行资源的隔离
         + LXC在资源管理方面依赖于Linux内核的cgroups子系统，cgroups子系统是Linux内核提供的一个基于进程组的资源管理的框架，可以为特定的进程组限定可以使用的资源
         + 运行环境的隔离使用了NameSpace，namespace隔离包括进程树，网络，用户id，以及挂载的文件系统，主要方式是在对应的隔离项上加tag
      + docker 后来移除了LXC，使用libcontainer（后来改名runc）
         + libcontainer 为 golang 实现，提供了虚拟化技术的api，可以简单理解为golang实现的LXC
   2. dockerfile 主要内容
      FROM: 原始镜像  
      剩下的基本类似于linux操作，比如拷贝代码，执行命令等
4. Prometheus && Grafana
   1. Prometheus 架构

       <img src="https://github.com/siyoo/devops/blob/master/images/prometheus.png" height="800px" width="800px" />

        Server 主要负责数据采集和存储，提供PromQL查询语言的支持  
        Alertmanager 警告管理器，用来进行报警  
        Push Gateway 支持临时性Job主动推送指标的中间网关

   2. Prometheus查询 promQL  
        参考 <a href="https://www.jianshu.com/p/93c840025f01" />简书-三无架构师-Prometheus</a>

   3. Grafana 支持 Graphite，Elasticsearch，InfluxDB，Prometheus，Cloudwatch，MySQL，OpenTSDB等时序性数据库的展示

## 其他
1. Gitops  
   GitOps这个术语是由Weaveworks的Alexis Richardson在一篇名为《Operation by Pull Request》的博文中创造的。其基本思想是通过向Git提交变更并使用Pull Request（以下简称PR）进行审批来管理Kubernetes上的资源
   * Jenkins X: 集成k8s的Jenkins (使用helm charts进行k8s集群的管理)
   * Argo CD:美观易用的gitops，详情见演示

2. 网关 && 负载均衡
   1. OpenResty: 利用 nginx 做服务器  
      OpenResty 是一组扩展Nginx功能的模块，开发人员可以使用 Lua 脚本语言调动 Nginx 支持的各种 C 以及 Lua 模块
      详情见demo
   2. Kong: Gateway + OpenResty  
      集成了网关功能的OpenResty

