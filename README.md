# 使用golang实现了mock.js功能(http://mockjs.com/examples.html)

#### 安装

go get github.com/liyanbing/go-mock

#### 使用
```cgo
package main

import "github.com/liyanbing/go-mock"

func main(){
   	cases := []struct {
   		Schema string
   	}{
   		{
   			Schema: schema,
   		},
   		{
   			Schema: "Mock.mock({\"name|1-9\": \"张\"})",
   		},
   		{
   			Schema: "Mock.mock({\"object|2\":{\"310000\":\"上海市\",\"320000\":\"江苏省\",\"330000\":\"浙江省\",\"340000\":\"安徽省\"}})",
   		},
   	}
   
   	for _, value := range cases {
   		m, err := go-mock.NewMock(value.Schema)
   		if err != nil {
   			t.Error(err)
   			return
   		}
   		fmt.Println(m.MockString(context.Background()))
   	}
}
```