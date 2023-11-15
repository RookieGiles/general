package response

type successStruct struct {
	Code int         `json:"code"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"data"`
}

type failureStruct struct {
	Code int         `json:"code"`
	Msg  interface{} `json:"msg"`
}

// Success 请求成功返回
func Success(data interface{}) *successStruct {
	json := &successStruct{Code: successCode, Msg: successMsg, Data: data}

	return json
}

// Failure 请求失败返回
func Failure(code int, data interface{}) *failureStruct {
	json := &failureStruct{Code: successCode, Msg: successMsg}

	return json
}
