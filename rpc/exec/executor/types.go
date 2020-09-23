// Code generated by goctl. DO NOT EDIT!
// Source: exec.proto

package executor

import (
		"errors"
		exec "github.com/weblfe/ffmpeg-service/rpc/exec/pb"
)

var errJsonConvert = errors.New("json convert error")

type (
	Any struct {
		TypeUrl string `json:"type_url,omitempty"`
		Value   []byte `json:"value,omitempty"`
	}

	Entry struct {
		Key   string `json:"key,omitempty"`
		Value *Any   `json:"value,omitempty"`
	}

	Error struct {
		Errmsg string `json:"errmsg,omitempty"`
		Errno  int32  `json:"errno,omitempty"`
	}

	MetaData struct {
		Data []*Entry `json:"data,omitempty"`
	}

	Request struct {
		Action     string   `json:"action,omitempty"`     // 命令
		Params     []string `json:"params,omitempty"`     // 参数
		ReturnBody string   `json:"returnBody,omitempty"` // 返回数据结构
		Input      string   `json:"input,omitempty"`      // 输入文件
	}

	Response struct {
		Msg   string        `json:"msg,omitempty"`
		Code  exec.Response_Code `json:"code,omitempty"`
		Data  []*Entry      `json:"data,omitempty"`
		Error *Error        `json:"error,omitempty"`
		Meta  *MetaData     `json:"meta,omitempty"`
	}
)
