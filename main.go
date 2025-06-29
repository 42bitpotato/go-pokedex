package main

type config struct {
	mapPage  int
	mapLimit int
}

func main() {
	cfg := &config{}
	startRepl(cfg)
}
