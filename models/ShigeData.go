package models

type ShigeDoc struct {
	title string
}

func TestData() ShigeDoc {
	data := ShigeDoc{title:"cestes"}
	return data
}

func TestString() string {
	return "这是测试"
}
