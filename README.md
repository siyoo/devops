## 整体架构
<!--![整体架构](https://github.com/siyoo/devops/blob/master/images/整体架构.png)-->
<img src="https://github.com/siyoo/devops/blob/master/images/整体架构.png" height="400px" width="400px" />

---

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

+ 如何评估我们是否需要微服务

[8-ways-to-test-your-readiness-for-microservices]:https://www.zdnet.com/article/8-ways-to-test-your-readiness-for-microservices/ "8-ways-to-test-your-readiness-for-microservices"


