# client go-micro 客户端
## client 客户端
```
import (
	"context"

	client "github.com/gomsa/tools/k8s/client"
)
err = client.Call(ctx, ServiceName, "Service.Method", req, res)
if err != nil {
	return err
}
```