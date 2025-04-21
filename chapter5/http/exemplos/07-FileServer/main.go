package main


func main() {
	mux := http.NewServerMux()

	fileServer := http.FileServer(http.Dir("./public"))
}