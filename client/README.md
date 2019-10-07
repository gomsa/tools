# client go-micro 客户端
## client 客户端
```
import (
    "context"
    client "github.com/gomsa/tools/client"
)

func main() {
    err = client.Call(context.TODO(), ServiceName, "Service.Method", req, res)
    if err != nil {
        return err
    }
}
```