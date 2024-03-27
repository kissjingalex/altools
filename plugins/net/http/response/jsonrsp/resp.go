package jsonrsp

import (
	"altools/plugins/errcode"
	cjson "altools/plugins/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	"net/http"
)

const (
	Lang       = "_ctxLang_"
	respCtxKey = "_respCtxKey"
)

type StatusCodeError struct {
	ErrCode int
	ErrMsg  map[string]string
}

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	TraceId string      `json:"traceId"`
}

func ApiError(c *gin.Context, status errcode.StatusCodeError, message ...string) {
	ApiResponse(c, status, nil, message...)
}

func ApiResponse(c *gin.Context, status errcode.StatusCodeError, data interface{}, messages ...string) {
	rsp := &Response{
		Status:  status.ErrCode,
		Data:    data,
		Message: status.Message(c.GetString(Lang)),
	}
	if len(messages) > 0 {
		rsp.Message = fmt.Sprintf(rsp.Message, messages)
	}

	if data == nil {
		rsp.Data = struct{}{}
	}

	bs, err := cjson.JSON.Marshal(rsp)
	if err != nil {
		fmt.Printf("error=%v\n", err)
	}
	c.Set(respCtxKey, string(bs))

	// ------------------------
	// ------------------------
	// 压缩(优先br，然后gzip， 最后不压缩)
	//if len(bs) > 256 {
	//	var (
	//		encodings  = c.GetHeader("Accept-Encoding")
	//		encoding   string
	//		buf        = new(bytes.Buffer)
	//		compressor io.WriteCloser
	//	)
	//	if strings.Contains(encodings, "br") {
	//		encoding = "br"
	//		compressor = brotli.NewWriterLevel(buf, brotli.DefaultCompression)
	//	} else if strings.Contains(encodings, "gzip") {
	//		encoding = "gzip"
	//		compressor, err = gzip.NewWriterLevel(buf, gzip.DefaultCompression)
	//		if err != nil {
	//			log.Error("gzipNewWriterLevel").Msgf("%v", err)
	//			compressor = nil
	//		}
	//	}
	//	if compressor != nil {
	//		_, errW := compressor.Write(bs)
	//		if errW != nil {
	//			log.Error("compressorWrite").Msgf("%v", err)
	//		}
	//		errClose := compressor.Close()
	//		if errClose != nil {
	//			log.Error(" compressorClose").Msgf("%v", err)
	//		}
	//
	//		if errW == nil && errClose == nil {
	//			c.Header("Content-Encoding", encoding)
	//			bs = buf.Bytes()
	//		}
	//	}
	//}
	// ------------------------
	// ------------------------

	c.Render(http.StatusOK, render.Data{
		ContentType: "application/json; charset=utf-8",
		Data:        bs,
	})
	c.Abort()
}

func GetCtxStoreRspStr(c *gin.Context) string {
	return c.GetString(respCtxKey)
}
