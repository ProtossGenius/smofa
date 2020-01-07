package ofalz

import (
	"strings"
	"fmt"
)

type LineCmdFunc func(this *OutFileAlz, line string) error

type OutFileAlz struct {
	CmdList    []string
	Ptr        int
	LineCmdMap map[rune]LineCmdFunc
	Ocm        map[string]RegitstCmd
	Kmap       map[string]string
	Kamap      map[string][]string
	TagMap     map[string]int
	IntMap     map[string]int
	FloatMap   map[string]float64
	MapMap     map[string]map[string]interface{}
}

func NewOfalz(ocm map[string]RegitstCmd) *OutFileAlz {
	return  &OutFileAlz{Ocm: ocm, Kmap: map[string]string{}, Kamap: map[string][]string{}, IntMap: map[string]int{},
		FloatMap: map[string]float64{},TagMap:map[string]int{}, MapMap: map[string]map[string]interface{}{}, LineCmdMap: map[rune]LineCmdFunc{}}
}

type RegitstCmd func(this *OutFileAlz, prms ...string) error

func (this *OutFileAlz) Run() error {
	for i := this.Ptr; i < len(this.CmdList); i++{
		rStr := []rune(this.CmdList[i])
		if len(rStr) > 0 && rStr[0] == ':'{
			this.TagMap[strings.TrimSpace(this.CmdList[i][1:])] = i
		}
	}
	for this.Ptr < len(this.CmdList) {
		this.Exec()
	}
	return nil
}

func (this *OutFileAlz) Exec() error {
	line := this.CmdList[this.Ptr]
	defer func() {this.Ptr++}()
	if len(line) == 0 {
		return nil
	}
	if lc, ok := this.LineCmdMap[[]rune(line)[0]]; ok {
		err := lc(this, line[1:])
		if err != nil {
			return err
		}
	}else {
		this.Warning("no cmd line : [%d]%s", this.Ptr, line)
	}
	return nil
}

func (this *OutFileAlz) AppendExec(cmd ...string) error {
	this.CmdList = append(this.CmdList, cmd...)
	return this.Run()
}

func (this *OutFileAlz) Warning(f string, v ...interface{}){
	fmt.Printf("[Warning]" + f + "\n", v...)
}

func (this *OutFileAlz) Print(f string, v ...interface{}){
	fmt.Printf(f + "\n", v...)
}

func (this *OutFileAlz) GetString(key string) (val string, ok bool) {
	if val, ok = this.Kmap[key]; ok {
		return
	} else if arr, ok := this.Kamap[key]; ok {
		return strings.Join(arr, "\n"), ok
	}
	return "", false
}

func (this *OutFileAlz) GetTag(tag string) int{
	tag = strings.TrimSpace(tag)
	if tag == "end"{
		return len(this.CmdList)
	}
	if val, ok := this.TagMap[tag]; ok{
		return val
	}
	return -1
}

func (this *OutFileAlz) Goto(tag string) {
	idx := this.GetTag(tag)
	if idx < 0{
		this.Warning("No such Tag %s", tag)
	}else {
		this.Ptr = idx
	}
}

func (this *OutFileAlz) Delete(key string) {
	delete(this.Kmap, key)
	delete(this.Kamap, key)
}

func (this *OutFileAlz) Put(key string, val ...string) {
	this.Delete(key)
	if len(val) == 0 {
		return
	}
	if len(val) == 1 {
		this.Kmap[key] = val[0]
	} else {
		this.Kamap[key] = val
	}
}

func (this *OutFileAlz) AddLineCmd(s rune, lf LineCmdFunc) {
	this.LineCmdMap[s] = lf
}
