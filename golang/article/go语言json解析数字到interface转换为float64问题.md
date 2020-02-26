#### 问题描述
json解析数字到interface, 默认是float64类型，当数字的整数部分超过6位，打印就会变成科学记数法。

##### 如下例子所示：
```go
package main
 
import (
   "fmt"
   "encoding/json"
)
 
func main() {
   var result map[string]interface{}
   jsonStr := `{"value" : 1234567}`
   json.Unmarshal([]byte(jsonStr), &result)
   fmt.Println(result["value"])
}
```
##### 打印结果：
1.234567e+06

#### 解决方案
1）字符串
  fmt.Sprintf("%.f", result["value"])
2）整数
  int(result["value"].(float64))