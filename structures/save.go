package structures

type Save struct {
	ID        uint64 `json:"id"`
	Username  string `json:"username"`
	Password  uint64 `json:"password"`
	Timestamp uint64 `json:"timestamp"`
	SaveData  string `json:"savedata"`
}

type Credentials struct {
	Username string `json:"username"`
	Password uint64 `json:"password"`
}

type OtherStorage struct {
	Content string `json:"content"`
}
