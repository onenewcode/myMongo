# 文档
# filter
## D
D 是 BSON 文档的有序表示形式。当元素的顺序很重要时，应使用此类型，
>bson.D{{"name", "Madame Vo"}}

结构
```go
type D []E
```
## A
A 是 BSON 数组的有序表示。
>bson.A{"bar", "world", 3.14159, bson.D{{"qux", 12345}}}

结构
```go
type A []interface{}
```
## E
E 表示 D 的 BSON 元素。它通常用于 D 内部。
结构
```go
type E struct {
	Key   string
	Value interface{}
}
```

## M
M 是 BSON 文档的无序表示。当元素的顺序不事。此类型在编码和解码时作为常规 map[string]interface{} 进行处理。