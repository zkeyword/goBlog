package result

import (
	"BLOG/config"
	"github.com/kataras/iris"
)

// Map iris
func Map(data map[string]interface{}) iris.Map {
	var resultMap = iris.Map{}

	for key, value := range data {
		if key == "Title" {
			value = value.(string) + " - " + config.SystemConfig.AppTitle
		}
		resultMap[key] = value
	}
	return resultMap
}
