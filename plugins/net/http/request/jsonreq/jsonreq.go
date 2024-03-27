package jsonreq

import (
	cjson "altools/plugins/json"
	"fmt"
	"github.com/gin-gonic/gin"
)

// JSONDecode json 解析
func JSONDecode(c *gin.Context, val interface{}) (err error) {
	if val == nil {
		return nil
	}
	raw, err := c.GetRawData()
	if err != nil {
		fmt.Printf("error=%v\n", err)
		return
	}

	if err = cjson.JSON.Unmarshal(raw, val); err != nil {
		fmt.Printf("error=%v\n", err)
	}

	return err
}
