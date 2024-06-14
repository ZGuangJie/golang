### GIN框架

​		web浏览器所有的交互都是一次**请求(request)**，一次**响应(response)**。

![](https://cdn.jsdelivr.net/gh/ZGuangJie/GoPicture/golang/202406141449428.png)

#### 1、gin框架安装与使用

​		安装gin框架：

```go
go get -u github.com/gin-gonic/gin
```

##### RESTful API:

​		REST与技术无关，代表的是一种软件架构风格，REST是Representational State Transfer的简称，中文翻译为“表征状态转移”或“表现层状态转化”。

推荐阅读[阮一峰 理解RESTful架构](http://www.ruanyifeng.com/blog/2011/09/restful.html)

简单来说，REST的含义就是客户端与Web服务器之间进行交互的时候，使用HTTP协议中的4个请求方法代表不同的动作。

- `GET`用来获取资源
- `POST`用来新建资源
- `PUT`用来更新资源
- `DELETE`用来删除资源。