package validate

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

const tagName = "validate"

//邮箱验证正则
var mailRe = regexp.MustCompile(`\A[\w+\-.]+@[a-z\d\-]+(\.[a-z]+)*\.[a-z]+\z`)

//验证接口
type Validator interface {
	Validate(interface{}) (bool, error)
}

type defaultValidator struct {
}

func (v defaultValidator) Validate(val interface{}) (bool, error) {
	return true, nil
}

type stringValidator struct {
	Min int
	Max int
}

func (v stringValidator) Validate(val interface{}) (bool, error) {
	l := len(val.(string))

	if l == 0 {
		return false, fmt.Errorf("不能为空")
	}

	if l < v.Min {
		return false, fmt.Errorf("字符不能少于:%v个", v.Min)
	}

	if v.Max >= v.Min && l > v.Max {
		return false, fmt.Errorf("字符不能长于:%v个", v.Max)
	}

	return true, nil
}

type float32Validator struct {
	Min float32
	Max float32
}

func (v float32Validator) Validate(val interface{}) (bool, error) {
	num := val.(float32)

	if num < v.Min {
		return false, fmt.Errorf("值不能小于:%v", v.Min)
	}

	if v.Max >= v.Min && num > v.Max {
		return false, fmt.Errorf("值不能大于:%v", v.Max)
	}

	return true, nil
}

type float64Validator struct {
	Min float64
	Max float64
}

func (v float64Validator) Validate(val interface{}) (bool, error) {
	num := val.(float64)

	if num < v.Min {
		return false, fmt.Errorf("值不能小于:%v", v.Min)
	}

	if v.Max >= v.Min && num > v.Max {
		return false, fmt.Errorf("值不能大于:%v", v.Max)
	}

	return true, nil
}

type uSignValidator struct {
	Min uint
	Max uint
}

func (v uSignValidator) Validate(val interface{}) (bool, error) {
	num := val.(uint)

	if num < v.Min {
		return false, fmt.Errorf("值不能小于:%v", v.Min)
	}

	if v.Max >= v.Min && num > v.Max {
		return false, fmt.Errorf("值不能大于:%v", v.Max)
	}

	return true, nil
}

type numberValidator struct {
	Min int
	Max int
}

func (v numberValidator) Validate(val interface{}) (bool, error) {
	num := val.(int)

	if num < v.Min {
		return false, fmt.Errorf("值不能小于:%v", v.Min)
	}

	if v.Max >= v.Min && num > v.Max {
		return false, fmt.Errorf("值不能大于:%v", v.Max)
	}

	return true, nil
}

type EmailValidator struct {
}

func (v EmailValidator) Validate(val interface{}) (bool, error) {
	if !mailRe.MatchString(val.(string)) {
		return false, fmt.Errorf("is not a valid email address")
	}
	return true, nil
}

func getValidatorFromTag(tag string) Validator {
	args := strings.Split(tag, ",")

	switch args[0] {
	case "number":
		validator := numberValidator{}
		//将structTag中的min和max解析到结构体中
		fmt.Sscanf(strings.Join(args[1:], ","), "min=%d,max=%d", &validator.Min, &validator.Max)
		return validator
	case "float32":
		validator := float32Validator{}
		//将structTag中的min和max解析到结构体中
		fmt.Sscanf(strings.Join(args[1:], ","), "min=%d,max=%d", &validator.Min, &validator.Max)
		return validator
	case "float64":
		validator := float64Validator{}
		//将structTag中的min和max解析到结构体中
		fmt.Sscanf(strings.Join(args[1:], ","), "min=%d,max=%d", &validator.Min, &validator.Max)
		return validator
	case "usign":
		validator := uSignValidator{}
		//将structTag中的min和max解析到结构体中
		fmt.Sscanf(strings.Join(args[1:], ","), "min=%d,max=%d", &validator.Min, &validator.Max)
		return validator
	case "string":
		validator := stringValidator{}
		fmt.Sscanf(strings.Join(args[1:], ","), "min=%d,max=%d", &validator.Min, &validator.Max)
		return validator
	case "email":
		return EmailValidator{}
	}

	return defaultValidator{}
}

func ValidateStruct(s interface{}) error {
	v := reflect.ValueOf(s)

	for i := 0; i < v.NumField(); i++ {
		//利用反射获取structTag
		tag := v.Type().Field(i).Tag.Get(tagName)

		if tag == "" || tag == "-" {
			continue
		}

		validator := getValidatorFromTag(tag)

		valid, err := validator.Validate(v.Field(i).Interface())
		if !valid && err != nil {
			return fmt.Errorf("%s %s", v.Type().Field(i).Name, err.Error())
		}
	}

	return nil
}
