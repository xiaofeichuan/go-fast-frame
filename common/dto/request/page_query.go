package request

//分页查询请求
type PageQuery struct {
	PageNum  int `json:"pageNum" form:"pageNum"`   // 页码
	PageSize int `json:"pageSize" form:"pageSize"` // 每页数量
}
