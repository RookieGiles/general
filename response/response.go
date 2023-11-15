package response

type SuccessStruct struct {
	Code int         `json:"code"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"data"`
}

type FailureStruct struct {
	Code int         `json:"code"`
	Msg  interface{} `json:"msg"`
}

// Success 请求成功返回
func Success(data interface{}) *SuccessStruct {
	json := &SuccessStruct{Code: SuccessCode, Msg: SuccessMsg, Data: data}

	return json
}

// Failure 请求失败返回
func Failure(code int, data interface{}) *FailureStruct {
	json := &FailureStruct{Code: SuccessCode, Msg: SuccessMsg}

	return json
}
