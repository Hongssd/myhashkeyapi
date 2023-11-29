package myhashkeyapi

import (
	"fmt"
	"strings"
)

type HashkeyErrorRes struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
}

type HashkeyTimeRes struct {
	InTime  string `json:"inTime"`  //REST网关接收请求时的时间戳，Unix时间戳的微秒数格式，如 1597026383085123返回的时间是请求验证后的时间
	OutTime string `json:"outTime"` //REST网关发送响应时的时间戳，Unix时间戳的微秒数格式，如 1597026383085123
}
type HashkeyRestRes[T any] struct {
	HashkeyErrorRes   //错误信息
	HashkeyTimeRes    //时间戳
	Result          T `json:"result"` //请求结果
}

func handlerCommonRest[T any](data []byte) (*HashkeyRestRes[T], error) {
	res := &HashkeyRestRes[T]{}
	var err error
	// log.Warn(string(data))
	if strings.Contains(string(data), "code") && !strings.HasPrefix(string(data), "[") {
		err = json.Unmarshal(data, res)
		if err != nil {
			log.Error("rest返回值获取失败", err)
		}
	} else {
		var result T
		err = json.Unmarshal(data, &result)
		if err != nil {
			log.Error("rest返回值序列化错误", err)
		}
		res.Result = result
	}
	return res, err
}
func (err *HashkeyErrorRes) handlerError() error {
	if err.Code != "" && err.Msg != "" {
		return fmt.Errorf("request error:[code:%v][message:%v]", err.Code, err.Msg)
	} else {
		return nil
	}

}
