package util

import (
	"encoding/json"
	"github.com/donkeywon/gtil/util"
	"net/http"
	"sync"
)

const (
	ResponseCodeOK                  = 0
	ResponseCode420                 = 420
	ResponseCodeNicknameExist       = 255
	ResponseCodeNicknameTooShort    = 1
	ResponseCodeNickNameTooShort256 = 256
	ResponseCode404                 = 404
)

var _responseMsgMap = map[int]string{
	ResponseCodeOK:                  "",
	ResponseCode420:                 "Please play as PMC and go through the offline settings screen before pressing ready.",
	ResponseCodeNicknameExist:       "The nickname is already in use",
	ResponseCodeNicknameTooShort:    "The nickname is too short",
	ResponseCodeNickNameTooShort256: "The nickname is too short",
	ResponseCode404:                 "UNHANDLED REQUEST",
}

func GetResponseMsg(code int) string {
	return _responseMsgMap[code]
}

type response struct {
	ErrCode int         `json:"err"`
	ErrMsg  string      `json:"errmsg"`
	Data    interface{} `json:"data"`
}

var (
	_pool = NewPool()
)

func GetResponseWrapper() *HttpResponse {
	hr := _pool.Get().(*HttpResponse)
	hr.r.ErrCode = ResponseCodeOK
	hr.r.ErrMsg = GetResponseMsg(hr.r.ErrCode)
	hr.r.Data = nil
	hr.p = _pool
	return hr
}

func GetResponseWrapperFromBytes(bs []byte) (*HttpResponse, error) {
	var resp response
	err := json.Unmarshal(bs, &resp)
	if err != nil {
		return nil, err
	}

	hr := _pool.Get().(*HttpResponse)
	hr.r.ErrCode = resp.ErrCode
	hr.r.ErrMsg = resp.ErrMsg
	hr.r.Data = resp.Data
	hr.p = _pool
	return hr, nil
}

func GetResponseWrapperFromString(s string) (*HttpResponse, error) {
	return GetResponseWrapperFromBytes(util.String2Bytes(s))
}

type HttpResponse struct {
	r *response
	p *sync.Pool
}

func (hr *HttpResponse) Free() {
	hr.p.Put(hr)
}

func (hr *HttpResponse) SetErrCode(code int) {
	hr.r.ErrCode = code
}

func (hr *HttpResponse) GetErrCode() int {
	return hr.r.ErrCode
}

func (hr *HttpResponse) SetErrMsg(m string) {
	hr.r.ErrMsg = m
}

func (hr *HttpResponse) GetErrMsg() string {
	return hr.r.ErrMsg
}

func (hr *HttpResponse) SetData(d interface{}) {
	hr.r.Data = d
}

func (hr *HttpResponse) GetData() interface{} {
	return hr.r.Data
}

func (hr *HttpResponse) GetResponse() *response {
	return hr.r
}

func NewPool() *sync.Pool {
	return &sync.Pool{
		New: func() interface{} {
			return &HttpResponse{
				r: &response{},
			}
		},
	}
}

func HTTPResponse(errCode int, errMsg string, data interface{}, httpCode int, w http.ResponseWriter) error {
	r := GetResponseWrapper()
	r.SetErrCode(errCode)
	r.SetErrMsg(errMsg)
	r.SetData(data)
	defer r.Free()

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(httpCode)

	ret, err := json.Marshal(r.GetResponse())
	if err != nil {
		return err
	}

	_, err = w.Write(ret)
	if err != nil {
		return err
	}

	return nil
}
