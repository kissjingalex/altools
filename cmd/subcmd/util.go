package subcmd

import (
	"encoding/json"
	"fmt"
	"github.com/fbsobreira/gotron-sdk/pkg/common"
)

func PrintAsJson(v any) {
	asJSON, _ := json.Marshal(v)
	fmt.Println(common.JSONPrettyFormat(string(asJSON)))
}
