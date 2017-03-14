package models

type ShigeDoc struct {
	Title string
}

func TestData() *ShigeDoc {
	data := &ShigeDoc{Title:"cestes"}
	return data
}

func TestString() string {
	return "这是测试"
}
