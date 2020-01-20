package business_errors

import errors "github.com/kataras/iris/v12/core/errgroup"

var (
	TopicTitleNotBeEmpty = errors.New("帖子标题不能为空")
)
