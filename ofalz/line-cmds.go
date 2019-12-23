package ofalz

import (
	"strings"
)

func DftLcOfalz(ocm map[string]RegitstCmd) *OutFileAlz {
	res := NewOfalz(ocm)
	res.AddLineCmd('?', lc_debugInfo)
	res.AddLineCmd('@', lc_updateVar)
	res.AddLineCmd('>', lc_mutiLineVar)
	res.AddLineCmd('~', lc_goto)
	res.AddLineCmd('$', lc_execMuti)
	res.AddLineCmd('!', lc_execLine)
	return res
}

//?
func lc_debugInfo(this *OutFileAlz, line string) error {
	slt := strings.Split(line, ",")
	for _, key := range slt {
		key = strings.TrimSpace(key)
		str, ok := this.GetString(key)
		if !ok {
			this.Print("<?>: key %s not exist", key)
		} else {
			this.Print("<?>: key %s = %s", key, str)
		}
	}
	return nil
}

//@
func lc_updateVar(this *OutFileAlz, line string) error {
	slt := strings.Split(line, ",")
	for _, keyVal := range slt {
		eqaPos := strings.Index(keyVal, "=")
		if eqaPos == -1{
			this.Warning("<@>Unexpected Input", keyVal)
			continue
		}
		key := strings.TrimSpace(keyVal[:eqaPos])
		val := strings.TrimSpace(keyVal[eqaPos+1:])
		this.Put(key, val)
	}
	return nil
}

//>
func lc_mutiLineVar(this *OutFileAlz, line string) error {
	key := strings.TrimSpace(line)
	arrList := []string{}
	this.Ptr++
	for this.Ptr < len(this.CmdList){
		firstChar := []rune(this.CmdList[this.Ptr])[0]
		if _, ok := this.LineCmdMap[firstChar]; ok{
			break
		}
		if firstChar == '\\'{
			arrList = append(arrList, this.CmdList[this.Ptr][1:])
		}else {
			arrList = append(arrList, this.CmdList[this.Ptr])
		}
		this.Ptr++
	}
	this.Ptr--
	this.Kamap[key] = arrList
	return nil
}

//~
func lc_goto(this *OutFileAlz, line string) error {
	this.Goto(line)
	return nil

}

//$
func lc_execMuti(this *OutFileAlz, line string) error {
	cmd := strings.Split(line, " ")[0]
	if f, ok := this.Ocm[cmd]; ok{
		vars := strings.Split(line[len(cmd):], ",")
		for i := range vars{
			vars[i] = strings.TrimSpace(vars[i])
		}
		return f(this, vars...)
	}
	this.Warning("Register Cmd Not Found %s", cmd)
	return nil
}

//!
func lc_execLine(this *OutFileAlz, line string) error {
	cmd := strings.Split(line, " ")[0]
	if f, ok := this.Ocm[cmd]; ok{
		return f(this, strings.TrimSpace(line[len(cmd):]))
	}
	this.Warning("Register Cmd Not Found %s", cmd)
	return nil
}

//:
func lc_tag(this *OutFileAlz, line string) error {
	tag := strings.TrimSpace(line)
	this.TagMap[tag] = this.Ptr
	return nil
}

//$
func lc_cache(this *OutFileAlz, line string) error {
	key := strings.TrimSpace(line)
	this.Kamap[key] = []string{}
	return nil
}