package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"runtime"

	"github.com/zserge/lorca"
)

func main() {
	cliArgs := os.Args[1:]
	args := []string{}
	if runtime.GOOS == "linux" {
		args = append(args, "--class=Lorca")
	}
	ui, err := lorca.New("", "", 800, 800, args...)
	if err != nil {
		log.Fatal(err)
	}
	defer ui.Close()

	bindFunctions(ui)

	fmt.Println(cliArgs)

	if cliArgs[0] == "--dev" {
		serveDev(ui)
	} else {
		serveProd(ui)
	}

	sigc := make(chan os.Signal)
	signal.Notify(sigc, os.Interrupt)
	select {
	case <-sigc:
	case <-ui.Done():
	}

	log.Println("exiting...")
}

func serveProd(ui lorca.UI) {
	// Start listener
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()

	// Serve 'dist' directory created when building your project
	go http.Serve(ln, http.FileServer(http.Dir("./frontend/dist/my-app")))

	ui.Load(fmt.Sprintf("http://%s", ln.Addr()))

	fmt.Println("Serving on " + fmt.Sprintf("http://%s", ln.Addr()))
}

func serveDev(ui lorca.UI) {
	ui.Load("http://localhost:4200")
}
