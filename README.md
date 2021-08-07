# xiaoshouyi
> 销售易CRM SDK (Golang)

## 特性

* 通过XOQL查询数据
* Access Token 自动管理和续期

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
