// Code generated by goctl. DO NOT EDIT.
package types

type ProduceReq struct {
	Msg string `form:"msg"`
}

type ProduceResp struct {
	Id int64 `json:"id"`
}