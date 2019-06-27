# go-json

```
go get github.com/aimuz/go-json
```
## 使用方法

和标准库一致，无需改动代码，通过编译参数自动选择 json 库

如使用 `jsoniter` 

```
go build -tags=jsoniter 
```

将会自动选择jsoniter作为json库
