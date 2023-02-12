package WeChatCustomerServiceSDK

import (
	"errors"
	"github.com/NICEXAI/WeChatCustomerServiceSDK/crypto"
)

// CryptoOptions 微信服务器验证参数
type CryptoOptions struct {
	Signature string `form:"msg_signature"`
	TimeStamp string `form:"timestamp"`
	Nonce     string `form:"nonce"`
	EchoStr   string `form:"echostr"`
}

// VerifyURL 验证请求参数是否合法
func (r *Client) VerifyURL(options CryptoOptions) (string, error) {
	wxCpt := crypto.NewWXBizMsgCrypt(r.token, r.encodingAESKey, r.corpID, crypto.XmlType)
	data, err := wxCpt.VerifyURL(options.Signature, options.TimeStamp, options.Nonce, options.EchoStr)
	if err != nil {
		return "", errors.New(err.ErrMsg)
	}
	return string(data), nil
}

// DecryptMsg 解密消息
func (r *Client) DecryptMsg(options CryptoOptions, postData []byte) ([]byte, error) {
	wxCpt := crypto.NewWXBizMsgCrypt(r.token, r.encodingAESKey, r.corpID, crypto.XmlType)
	message, status := wxCpt.DecryptMsg(options.Signature, options.TimeStamp, options.Nonce, postData)
	if status != nil && status.ErrCode != 0 {
		return nil, errors.New(status.ErrMsg)
	}
	return message, nil
}

// 加密消息
func (r *Client) EncryptMsg(options CryptoOptions, postData []byte) ([]byte, error) {
	wxCpt := crypto.NewWXBizMsgCrypt(r.token, r.encodingAESKey, r.corpID, crypto.XmlType)
	message, status := wxCpt.EncryptMsg(string(postData), options.TimeStamp, options.Nonce)
	if status != nil && status.ErrCode != 0 {
		return nil, errors.New(status.ErrMsg)
	}
	return message, nil
}
