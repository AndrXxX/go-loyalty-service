package entities

type User struct {
	Login    string `json:"login" valid:"minstringlength(3)"`
	Password string `json:"password" valid:"minstringlength(5)"`
}
