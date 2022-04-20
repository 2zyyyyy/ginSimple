package trans

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_trans "github.com/go-playground/validator/v10/translations/zh"
)

// 翻译器封装

type Trans struct {
	uni      *ut.UniversalTranslator
	validate *validator.Validate
	trans    ut.Translator
}

// NewTrans 创建翻译器
func NewTrans() (*Trans, error) {
	// 初始化翻译器
	tran := new(Trans)
	zhtr := zh.New()
	tran.uni = ut.New(zhtr, zhtr)
	tran.trans, _ = tran.uni.GetTranslator("zh")

	// 创建 validator
	tran.validate = binding.Validator.Engine().(*validator.Validate)

	// 注册翻译器
	if err := zh_trans.RegisterDefaultTranslations(tran.validate, tran.trans); err != nil {
		return nil, err
	}
	return tran, nil
}

// Translate 翻译错误信息
func (tran *Trans) Translate(err error) map[string][]string {
	var result = make(map[string][]string)
	errors := err.(validator.ValidationErrors)
	for _, err := range errors {
		result[err.Field()] = append(result[err.Field()], err.Translate(tran.trans))
	}
	return result
}
