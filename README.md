
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

2. 

