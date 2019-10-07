# client go-micro 客户端
## client 客户端
```
import (
    "context"
    client "github.com/gomsa/tools/client"
)

func main() {
    err = client.Call(ctx, ServiceName, "Service.Method", req, res)
    if err != nil {
        return err
    }
}
```