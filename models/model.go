package model

type Task struct {
	Id          int    `json:"id"`
	Task        string `json:"task"`
	Description string `json:"description"`
	UserId      int    `json:"user_id"`
}

type User struct {
	UserId   int     `json:"user_id"`
	UserName string  `json:"user_name"`
	Age      int     `json:"age"`
	Address  Address `json:"address"`
}

type Address struct {
	Address1 string `json:"address1"`
	Address2 string `json:"address2"`
	Zip      int32  `json:"zip"`
	City     string `json:"city"`
	State    string `json:"state"`
}
