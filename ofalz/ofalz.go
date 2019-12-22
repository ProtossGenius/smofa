package ofalz

import "strings"

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
	res := &OutFileAlz{Ocm: ocm, Kmap: map[string]string{}, Kamap: map[string][]string{}, IntMap: map[string]int{},
		FloatMap: map[string]float64{}, MapMap: map[string]map[string]interface{}{}, LineCmdMap: map[rune]LineCmdFunc{}}
	return res
}

type RegitstCmd func(this *OutFileAlz, prms ...string) error

func (this *OutFileAlz) Run() error {
	for this.Ptr < len(this.CmdList) {
		line := this.CmdList[this.Ptr]
		if len(line) == 0 {
			continue
		}
		if lc, ok := this.LineCmdMap[[]rune(line)[0]]; ok {
			err := lc(this, line)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (this *OutFileAlz) Get(key string) (val string, ok bool) {
	if val, ok = this.Kmap[key]; ok {
		return
	} else if arr, ok := this.Kamap[key]; ok {
		return strings.Join(arr, "\n"), ok
	}
	return "", false
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
