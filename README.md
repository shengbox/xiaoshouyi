# xiaoshouyi
> 销售易CRM SDK (Golang)

## 特性

* 支持第三方应用提供商的应用套件相关接口
* 企业号相关的所有 API 同时支持基于应用套件级别的调用和企业号单独调用
* 支持企业号最新的异步任务 API
* Access Token 自动管理和续期
* 支持 Access Token 超期或失效导致接口调用错误时重新获取并自动重试一次当前调用的 API
* 提供被动接收消息（事件）的解析方法，以及生成被动响应消息的方法

## 安装
```bash
$ go get -u github.com/shengbox/xiaoshouyi
```


## 使用

```go
import "github.com/shengbox/xiaoshouyi/api"

type User struct {
    api.User
    DimDepart int64 `json:"dimDepart"`
}

func main() {
    api.Debug = true
    xsy := api.New(clientId, clientSecret, username, password, redirectUri)
    var result []api.ActivityRecord
    _ = xsy.XOQLRecords(2, api.ActivityRecord{}, nil, &result)
    for _, it := range result {
        fmt.Println(it)
    }
    
    // 扩展对象
    var users []User
    _ = xsy.XOQLRecords(1, User{}, "WHERE status = 1", &users)
    for _, it := range users {
        fmt.Println(it)
    }
}

```
