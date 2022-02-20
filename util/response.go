package util

import (
	"bytes"
	"compress/zlib"
	"encoding/json"
	jsonvalue "github.com/Andrew-M-C/go.jsonvalue"
	"github.com/donkeywon/gtil/util"
	"io"
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

	HeaderKeyContentType = "Content-type"
	ContentTypeJson      = "application/json"
)

var _httpResponseMsgMap = map[int]string{
	ResponseCodeOK:                  "",
	ResponseCode420:                 "Please play as PMC and go through the offline settings screen before pressing ready.",
	ResponseCodeNicknameExist:       "The nickname is already in use",
	ResponseCodeNicknameTooShort:    "The nickname is too short",
	ResponseCodeNickNameTooShort256: "The nickname is too short",
	ResponseCode404:                 "UNHANDLED REQUEST",
}

func GetResponseMsg(code int) string {
	return _httpResponseMsgMap[code]
}

type httpResponse struct {
	ErrCode int         `json:"err"`
	ErrMsg  string      `json:"errmsg"`
	Data    interface{} `json:"data"`
}

var (
	_pool = NewPool()
)

func GetResponseWrapper() *httpResponseWrapper {
	hr := _pool.Get().(*httpResponseWrapper)
	hr.r.ErrCode = ResponseCodeOK
	hr.r.ErrMsg = GetResponseMsg(hr.r.ErrCode)
	hr.r.Data = nil
	hr.p = _pool
	return hr
}

func GetResponseWrapperFromBytes(bs []byte) (*httpResponseWrapper, error) {
	var resp httpResponse
	err := json.Unmarshal(bs, &resp)
	if err != nil {
		return nil, err
	}

	hr := _pool.Get().(*httpResponseWrapper)
	hr.r.ErrCode = resp.ErrCode
	hr.r.ErrMsg = resp.ErrMsg
	hr.r.Data = resp.Data
	hr.p = _pool
	return hr, nil
}

func GetResponseWrapperFromString(s string) (*httpResponseWrapper, error) {
	return GetResponseWrapperFromBytes(util.String2Bytes(s))
}

type httpResponseWrapper struct {
	r *httpResponse
	p *sync.Pool
}

func (hr *httpResponseWrapper) Free() {
	hr.p.Put(hr)
}

func (hr *httpResponseWrapper) SetErrCode(code int) {
	hr.r.ErrCode = code
}

func (hr *httpResponseWrapper) GetErrCode() int {
	return hr.r.ErrCode
}

func (hr *httpResponseWrapper) SetErrMsg(m string) {
	hr.r.ErrMsg = m
}

func (hr *httpResponseWrapper) GetErrMsg() string {
	return hr.r.ErrMsg
}

func (hr *httpResponseWrapper) SetData(d interface{}) {
	hr.r.Data = d
}

func (hr *httpResponseWrapper) GetData() interface{} {
	return hr.r.Data
}

func (hr *httpResponseWrapper) GetResponse() *httpResponse {
	return hr.r
}

func NewPool() *sync.Pool {
	return &sync.Pool{
		New: func() interface{} {
			return &httpResponseWrapper{
				r: &httpResponse{},
			}
		},
	}
}

func DoResponse(errCode int, errMsg string, data interface{}, w http.ResponseWriter) error {
	r := GetResponseWrapper()
	r.SetErrCode(errCode)
	r.SetErrMsg(errMsg)
	r.SetData(data)
	defer r.Free()

	return DoResponseJson(r.GetResponse(), w)
}

func DoResponseJson(data interface{}, w http.ResponseWriter) error {
	enc, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return DoResponseJsonBytes(enc, w)
}

func DoResponseJsonValue(v *jsonvalue.V, w http.ResponseWriter) error {
	enc, err := v.Marshal()
	if err != nil {
		return err
	}

	return DoResponseJsonBytes(enc, w)
}

func DoResponseJsonString(data string, w http.ResponseWriter) error {
	return DoResponseJsonBytes(util.String2Bytes(data), w)
}

func DoResponseJsonBytes(data []byte, w http.ResponseWriter) error {
	w.Header().Set(HeaderKeyContentType, ContentTypeJson)

	return DoResponseDirect(data, http.StatusOK, w)
}

func DoResponseZlibJson(data interface{}, w http.ResponseWriter) error {
	j, err := json.Marshal(data)
	if err != nil {
		return err
	}

	zr, err := zlib.NewReader(bytes.NewReader(j))
	if err != nil {
		return err
	}

	var out bytes.Buffer
	_, err = io.Copy(&out, zr)
	if err != nil {
		return err
	}

	return DoResponseJsonBytes(out.Bytes(), w)
}

func DoResponseDirect(data []byte, httpCode int, w http.ResponseWriter) error {
	w.WriteHeader(httpCode)
	_, err := w.Write(data)
	return err
}
