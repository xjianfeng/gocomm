package lxml

import (
	"bytes"
	"encoding/xml"
	"fmt"
)

func parse_token(decoder *xml.Decoder, tt xml.Token, ret map[string]string) {
	token := tt.(xml.StartElement)
	name := token.Name.Local

	var t xml.Token
	var err error
	for t, err = decoder.Token(); err == nil; t, err = decoder.Token() {
		switch t.(type) {
		// 处理元素开始（标签）
		case xml.StartElement:
			parse_token(decoder, t, ret)
			// 处理元素结束（标签）
		case xml.EndElement:
			return
		case xml.CharData:
			if name == "xml" {
				continue
			}
			ret[name] = fmt.Sprintf("%s", t)
		default:
			return
		}
	}
}

func ParseXml(xmlData []byte) (map[string]string, error) {
	var t xml.Token
	xmlData = bytes.TrimSpace(xmlData)
	xmlReader := bytes.NewReader(xmlData)
	decoder := xml.NewDecoder(xmlReader)

	t, decerr := decoder.Token()
	if decerr != nil {
		return nil, decerr
	}
	ret := make(map[string]string)
	parse_token(decoder, t, ret)
	return ret, nil
}
