package hello

func Hello() string {
	return "Hello World"
}

func HelloName(name string) string{
	if name==""{
		return "Hello, World"
	}
	return "Hello, "+name
}