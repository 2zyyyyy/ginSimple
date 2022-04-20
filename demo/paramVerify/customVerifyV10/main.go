package main

import (
	"fmt"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"reflect"
	"strings"
)

type User struct {
	ID     int64  `json:"id" validate:"gt=0"`
	Name   string `json:"name" validate:"required"`
	Gender string `json:"gender" validate:"required,oneof=man woman"`
	Age    uint8  `json:"age" validate:"required,gte=0,lte=130"`
	Email  string `json:"email" validate:"required,email"`
}

func main() {
	validate := validator.New()
	user := &User{
		ID:     1001,
		Name:   "wanli",
		Gender: "boy",
		Age:    10,
		Email:  "golang@google.gamil.com",
	}
	// 错误信息翻译成中文
	zh := zh.New()
	uni := ut.New(zh)
	trans, _ := uni.GetTranslator("zh")
	_ = en_translations.RegisterDefaultTranslations(validate, trans)

	// 注册一个函数，获取结构体字段的备用名称
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
		//fmt.Println(names)
		if name == "-" {
			return "j"
		}
		return name
	})
	err := validate.Struct(user)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		fmt.Println(validationErrors)
		return
	}
}
