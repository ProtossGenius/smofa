package main

import (
	"flag"
	"io/ioutil"
	"os"
	"strings"

	"github.com/ProtossGenius/smofa"
)

func FileReadAll(path string) ([]byte, error) {
	cfg, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	b, e := ioutil.ReadAll(cfg)
	cfg.Close()
	return b, e
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func rc_echo(this *smofa.OutFileAlz, prms ...string) error {
	this.Print("%v", prms)
	return nil
}

func main() {
	in := flag.String("in", "./datas/test.ofa", "in put file")
	flag.Parse()
	ofa := smofa.DftLcOfalz(map[string]smofa.RegistCmd{"echo": rc_echo})
	bts, err := FileReadAll(*in)
	check(err)
	ofa.AppendExec(strings.Split(string(bts), "\n")...)

}
