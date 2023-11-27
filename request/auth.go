package request

type Key struct {
	ClientID     string
	ClientSecret string
}

type Auth struct {
	Scope        string `form:"scope" binding:"required"`
	GrantType    string `form:"grant_type" binding:"required"`
	ClientID     string `form:"client_id" binding:"required"`
	ClientSecret string `form:"client_secret" binding:"required"`
	Username     string `form:"username" binding:"required"`
	Password     string `form:"password" binding:"required"`
}
