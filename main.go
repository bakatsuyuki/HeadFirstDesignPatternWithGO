package main

func main() {
	var hello = SingletonInstance()
	var hello2 = SingletonInstance()
	
	println(hello.toString())
	println(hello2.toString())
}
