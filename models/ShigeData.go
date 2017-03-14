package models

type ShigeDoc struct {
	Title string
	Author string
	Period string
	Doc string
	Labels string
}

func TestData() *ShigeDoc {
	data := &ShigeDoc{
		Title:"一剪梅",
		Author:"李清照",
		Period:"宋",
		Doc: "红藕香残玉簟秋\n轻解罗裳\n独上兰舟\n云中谁寄锦书来\n雁字回时\n月满西楼\n\n花自飘零水自流\n一种相思\n两处闲愁\n此情无计可消\n却上心头",
		Labels:"李清照#|#宋词"}
	return data
}

func TestString() string {
	return "这是测试"
}
