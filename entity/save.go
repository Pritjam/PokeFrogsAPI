package entity

type Save struct {
	ID        uint64 `json:"id"`
	Username  string `json:"username"`
	Password  uint64 `json:"password"`
	Timestamp uint64 `json:"timestamp"`
	SaveData  string `json:"savedata"`
}

type IDPass struct {
	ID       uint64 `json:"id"`
	Password uint64 `json:"password"`
}
