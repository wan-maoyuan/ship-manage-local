package proto

type LoginReq struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

type LoginResp struct {
	UserID      int64    `json:"user_id"`
	Account     string   `json:"account"`
	Password    string   `json:"password"`
	RoleID      int64    `json:"role_id"`
	Role        string   `json:"role"`
	Permissions []string `json:"permissions"`
}
