package helper

var version string = "1.0.0" //private access modifier
var Application string = "Golang" //public variable

func SayHelloHelper(name string) string { //public functions
	return "Hello " + name
}

func sayHello(name string) string  { // private function
	return "Hello " + Application	
}