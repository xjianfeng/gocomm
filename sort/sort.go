package sort

import (
	"bytes"
	"fmt"
	"sort"
)

type SortMap map[interface{}]interface{}

func (mp SortMap) CheckMapVaiable() bool {
	if len(mp) > 0 {
		return true
	}
	return false
}

func (mp SortMap) SortMapStrKey(desc bool) []string {
	if !mp.CheckMapVaiable() {
		return []string{}
	}
	keyList := []string{}
	for k, _ := range mp {
		keyList = append(keyList, k.(string))
	}
	if desc {
		sort.Sort(sort.Reverse(sort.StringSlice(keyList)))
	} else {
		sort.Strings(keyList)
	}
	return keyList
}

func (mp SortMap) SortMapIntKey(desc bool) []int {
	if !mp.CheckMapVaiable() {
		return []int{}
	}

	keyList := []int{}
	for k, _ := range mp {
		keyList = append(keyList, k.(int))
	}
	if desc {
		sort.Sort(sort.Reverse(sort.IntSlice(keyList)))
	} else {
		sort.Ints(keyList)
	}
	return keyList
}

func (mp SortMap) SortStrKeyUrlEncoded(desc bool) string {
	if !mp.CheckMapVaiable() {
		return ""
	}
	data := mp.SortMapStrKey(desc)
	buff := bytes.NewBufferString("")
	for _, k := range data {
		buff.WriteString(k)
		buff.WriteString("=")
		v := fmt.Sprintf("%v", mp[k])
		buff.WriteString(v)
		buff.WriteString("&")
	}
	s := buff.String()
	return s[:len(s)-1]
}
