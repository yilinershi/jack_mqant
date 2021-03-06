# 目标
* 通过实际案例学习研究mqant
* 截止2021.10.22，项目仍在业余时间更新中

# 服务器结构
* SV_DB为数据库服务器，只接收其它服的rpc消息，因为是通过nats作为消息队列，所有的数据获取及修改都只能用该服务器去完成
* SV_Gate为默认的网关服务器，里面封装了自定义的agent,并将客户端的通信类型进行了分类
> 在SV_Gate中，自定义了agent，并重新定义了路由，完整的路由例如：SV_Lobby/Call_Auth
> 
> 路由由2层结构组成，第1层为服务器类型，如SV_Lobby,指客户端要的该服务由SV_Lobby提供
> 
> 第2层路由为服务方法，如Call_Auth,该路由在SV_Lobby中被监听，其会响应callAuth方法
> 
> 第2层路由中，如果类型为Call，即服务器必需作为响应,可直接回复session，即req-resp结构,如为Notify即为req-noResp结构
> 
> 对于服务器而言，到时候还会有Push类型，即noReq-resp结构，即服务器推送消息给客户端
* SV_Http为项目的http服务器，提供登录、注册、充值等相关服务
* SV_Lobby为大厅服务器，提供大厅基本功能，如数据展示，排行榜，创建房间等服务
* 所有服务器间进行rpc通信，rpc支持protobuf结构（mqant提供），如"SV_Lobby/Call_OnAuth"路由示例中提供了protobuf结构的rpc通信完整示例

# 登录流程
* SV_Http中提供了注册及登录的方法，这两个接口均为http接口
* 为什么使用Http（短连接）作登录注册而不使用TCP/WebSocket(长连接)去作登录及注册？
> 登录及注册功能天生是无状态的，如果用长连接登录失败，再去断开长连接，这种做法非常不可取
> 
> 服务器的长连接地址轻易暴露在客户端，在客户端被解包后会有大量风险
> 
> 目前的做法是，一个无状态的短连接（http）在握手时，动态的下发长连接的地址，对于存在风险的请求连接，可以下发带有欺骗性的其它地址防止被攻击
> 
> 这样做到长连接地址不被暴露，同时，服务器可在变更地址后或是应对特殊需求时，下发不同的地址。比如对于长连接地址被攻击，对于一旦捕获到攻击方ip，只需要相应的转发baidu.com为长连接地址，让他们转而攻击百度即可。
>
> 项目中，如下3个方法演示了登录注册的流程
```
http.HandleFunc("/handshake", this.handshake)  
http.HandleFunc("/login", this.login)
http.HandleFunc("/register", this.register)
SV_Lobby/HD_OnAuth/Call
```

> step1:handshake,客户端req中包括：客户端版本号，客户端渠道等信息，服务器在resp中下发：长连接地址，热更新地址，登录地址，注册地址等
> step2:注册及登录，登录后，服务器返回token给客户端
> step3:客户端进行长连接，第一个长连接消息为HD_OnAuth,用于验证长连接是否已授权过，即token是否正确。
> 以上3步后整个登录流程完成

