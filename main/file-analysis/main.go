package main

import (
	"flag"
	"github.com/ProtossGenius/smofa/ofalz"
	"os"
	"io/ioutil"
	"strings"
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
	if err != nil{
		panic(err)
	}
}

func rc_echo(this *ofalz.OutFileAlz, prms ...string) error{
	this.Print("%v", prms)
	return nil
}

func main() {
	in := flag.String("in", "./datas/test.ofa", "in put file")
	flag.Parse()
	ofa := ofalz.DftLcOfalz(map[string]ofalz.RegitstCmd{"echo": rc_echo})
	bts , err := FileReadAll(*in)
	check(err)
	ofa.AppendExec(strings.Split(string(bts), "\n")...)

}
