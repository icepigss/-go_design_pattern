package main

func Print(str string) {
	println("str:" + str)
}

func decoratorAddSuffix(fn func(str string)) func(str string) {
	return func(str string) {
		println("decorator before")
		str += "_suffix"
		fn(str)
		println("decorator after")
	}
}

func main() {
	decoratorAddSuffix(Print)("icepigss")
}
