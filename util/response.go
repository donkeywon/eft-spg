package util

import (
	"bytes"
	"compress/zlib"
	"encoding/json"
	"github.com/bytedance/sonic"
	"github.com/bytedance/sonic/ast"
	"github.com/donkeywon/gtil/util"
	"github.com/pkg/errors"
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

var (
	_pool = NewPool()
)

func GetResponseMsg(code int) string {
	return _httpResponseMsgMap[code]
}

type httpResponse struct {
	ErrCode int         `json:"err"`
	ErrMsg  string      `json:"errmsg"`
	Data    interface{} `json:"data"`
}

func GetResponseWrapperFromData(data interface{}) *httpResponseWrapper {
	return GetResponseWrapperFromInfo(ResponseCodeOK, "", data)
}

func GetResponseWrapperFromInfo(code int, msg string, data interface{}) *httpResponseWrapper {
	hr := _pool.Get().(*httpResponseWrapper)
	hr.r.ErrCode = code
	hr.r.ErrMsg = msg
	hr.r.Data = data
	hr.p = _pool
	return hr
}

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
	bs, err := encodeResp(data)
	if err != nil {
		return errors.Wrap(err, ErrResponse)
	}

	return DoResponseJsonBytes(bs, w)
}

func DoResponseJsonValue(v *ast.Node, w http.ResponseWriter) error {
	enc, err := v.MarshalJSON()
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

	return DoResponseBytes(data, http.StatusOK, w)
}

func DoResponseZlibJson(data interface{}, w http.ResponseWriter) error {
	j, err := encodeResp(data)
	if err != nil {
		return errors.Wrap(err, ErrResponse)
	}

	zr, err := zlib.NewReader(bytes.NewReader(j))
	if err != nil {
		return err
	}

	out := GetBuffer()
	defer out.Free()
	_, err = io.Copy(out, zr)
	if err != nil {
		return err
	}

	return DoResponseJsonBytes(out.Bytes(), w)
}

func DoResponseBytes(data []byte, httpCode int, w http.ResponseWriter) error {
	w.WriteHeader(httpCode)
	_, err := w.Write(data)
	return err
}

func DoResponseString(data string, httpCode int, w http.ResponseWriter) error {
	return DoResponseBytes(util.String2Bytes(data), httpCode, w)
}

func encodeResp(data interface{}) ([]byte, error) {
	var err error
	var bs []byte

	switch data.(type) {
	case []byte:
		bs = data.([]byte)
	case string:
		bs = util.String2Bytes(data.(string))
	case *ast.Node:
		if data.(*ast.Node).Exists() {
			bs, err = data.(*ast.Node).MarshalJSON()
			if err != nil {
				return nil, errors.Wrap(err, ErrEncodeResp)
			}
		}
	case ast.Node:
		d := data.(ast.Node)
		if d.Exists() {
			bs, err = d.MarshalJSON()
			if err != nil {
				return nil, errors.Wrap(err, ErrEncodeResp)
			}
		}
	case *httpResponseWrapper:
		d := data.(*httpResponseWrapper)
		defer d.Free()
		bs, err = sonic.Marshal(d.r)
		if err != nil {
			return nil, errors.Wrap(err, ErrEncodeResp)
		}
	default:
		bs, err = sonic.Marshal(data)
		if err != nil {
			return nil, errors.Wrap(err, ErrEncodeResp)
		}
	}

	return bs, nil
}
