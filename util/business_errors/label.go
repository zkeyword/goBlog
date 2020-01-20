package business_errors

import errors "github.com/kataras/iris/v12/core/errgroup"

var (
	LabelNameAlreadyExists = errors.New("用户名已存在")
	LabelNameNotExist      = errors.New("用户名不存在")
	LabelNameNotBeEmpty    = errors.New("用户名不能为空")
)
