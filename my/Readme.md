Gin 是在 net/http 上进行封装，提供feature。其主要优化点为：
- 使用最多九颗前缀树（每一个HTTPMethod（POST/GET/...）各一颗）进行路由，树的节点按照URL中的/进行层级划分。每个节点都会挂接若干请求处理函数构成一个处理请求链。当一个请求到来时，在这棵树上找到请求URL对应的节点，拿到对应的请求处理链来执行就完成了请求的处理。
- gin.Context   Context贯穿一个http请求的所有流程，包含全部上下文信息。net/http 解析 HTTP 请求每次生成新的 *http.Request 和 http.ResponseWriter；gin 解析 HTTP 数据到 *gin.Context，然后使用 sync.Pool 复用结构实例，减少对象的数量