// AUTOGENERATED - DO NOT EDIT

package client

import (
	"encoding/json"
	"fmt"

	"github.com/Arman92/go-tdlib/v2/tdlib"
)

// GetMessageThread Returns information about a message thread. Can be used only if message.can_get_message_thread == true
// @param chatID Chat identifier
// @param messageID Identifier of the message
func (client *Client) GetMessageThread(chatID int64, messageID int64) (*tdlib.MessageThreadInfo, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type":      "getMessageThread",
		"chat_id":    chatID,
		"message_id": messageID,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var messageThreadInfo tdlib.MessageThreadInfo
	err = json.Unmarshal(result.Raw, &messageThreadInfo)
	return &messageThreadInfo, err

}
