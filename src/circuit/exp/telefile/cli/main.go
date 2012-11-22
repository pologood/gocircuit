package main

import (
	"circuit/kit/tele/file"
	"circuit/exp/telefile/srv"

	_ "circuit/load"
	"circuit/use/circuit"
	"circuit/use/n"
	"io"
	"os"
)

func main() {
	println("Starting")
	r, _, err := circuit.Spawn(n.ParseHost("localhost"), []string{"/telefile"}, srv.App{}, "/tmp/telehelo")
	if err != nil {
		println("Oh oh", err.Error())
		return
	}
	fcli := file.NewFileClient(r[0].(circuit.X))
	defer func() {
		recover()
	}()
	io.Copy(os.Stdout, fcli)
}
