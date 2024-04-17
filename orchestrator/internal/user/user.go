package user

type User struct {
	ID       string `xml:"id" json:"id"`
	Username string `xml:"username" json:"username"`
	Password string `xml:"password" json:"password,omitempty"`
	IsAdmin  bool   `xml:"is_admin" json:"is_admin"`
}
