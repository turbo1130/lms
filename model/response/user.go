package response

type User struct {
	Id            string `json:"id"`
	Username      string `json:"username"`
	RoleEn        string `json:"roleEn"`
	Authorization string `json:"authorization"`
}
