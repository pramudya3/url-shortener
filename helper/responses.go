package helper

type Error struct {
	Field interface{} `json:"field"`
	Msg   interface{} `json:"msg"`
}

type Success struct {
	Data interface{} `json:"data"`
}
