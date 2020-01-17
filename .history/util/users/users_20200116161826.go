package users

// import (
// 	"BLOG/config"
// 	"BLOG/model"
// 	"BLOG/services"
// 	"github.com/kataras/iris/sessions"
// )

// // 获取当前登录的用户ID
// func GetCurrentUserID(sess *sessions.Session) uint {
// 	userID := sess.GetIntDefault(conf.SystemConfig.UserIDKey, 0)
// 	return uint(userID)
// }

// // 获取当前登录的用户
// func GetCurrentUser(sess *sessions.Session) *model.User {
// 	user, _ := userService.FindByID(GetCurrentUserID(sess))
// 	return user
// }
