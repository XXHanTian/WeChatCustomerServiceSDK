package WeChatCustomerServiceSDK

import (
	"encoding/xml"
	"github.com/NICEXAI/WeChatCustomerServiceSDK/util"
	"net/http"
	"strconv"
	"time"
)

func (c *Client) Response(writer http.ResponseWriter, request *http.Request, reply interface{}) {
	var err error
	output := []byte("") // 默认回复
	if reply != nil {
		output, err = xml.Marshal(reply)
		if err != nil {
			return
		}
		// 加密
		output, err = c.encryptReplyMessage(output)
		if err != nil {
			return
		}
	}

	_, err = writer.Write(output)
	return
}

// encryptReplyMessage 加密回复消息
func (c *Client) encryptReplyMessage(rawXmlMsg []byte) (replyEncryptMessage []byte, err error) {
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	nonce := util.GetRandString(6)
	option := CryptoOptions{
		TimeStamp: timestamp,
		Nonce:     nonce,
	}
	cipherText, err := c.EncryptMsg(option, rawXmlMsg)
	if err != nil {
		return
	}

	return cipherText, err
}
