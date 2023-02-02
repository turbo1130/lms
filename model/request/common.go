package request

// PageInfo Paging initialize input parameter structure
type PageReq struct {
	Page     int         `json:"page" form:"page"`         // 页码
	PageSize int         `json:"pageSize" form:"pageSize"` // 每页大小
	Keyword  string      `json:"keyword" form:"keyword"`   //关键字
	Data     interface{} `json:"data"`
}

// GetById Find by id structure
type GetById struct {
	ID int `json:"id" form:"id"` // 主键ID
}

func (r *GetById) Uint() uint {
	return uint(r.ID)
}

type IdsReq struct {
	Ids []int `json:"ids" form:"ids"`
}

// GetroleId Get role by id structure
type GetRoleId struct {
	RoleId uint `json:"roleId" form:"roleId"` // 角色ID
}

type Empty struct{}
