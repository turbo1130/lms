package response

type PageResult struct {
	Page     int         `json:"page"`
	PageSize int64       `json:"pageSize"`
	Total    int64       `json:"total"`
	List     interface{} `json:"list"`
}
