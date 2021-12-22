// AUTOGENERATED - DO NOT EDIT

package client

import (
	"encoding/json"
	"fmt"

	"github.com/Arman92/go-tdlib/v2/tdlib"
)

// GetMessageLink Returns an HTTPS link to a message in a chat. Available only for already sent messages in supergroups and channels. This is an offline request
// @param chatID Identifier of the chat to which the message belongs
// @param messageID Identifier of the message
// @param forAlbum Pass true to create a link for the whole media album
// @param forComment Pass true to create a link to the message as a channel post comment, or from a message thread
func (client *Client) GetMessageLink(chatID int64, messageID int64, forAlbum bool, forComment bool) (*tdlib.MessageLink, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type":       "getMessageLink",
		"chat_id":     chatID,
		"message_id":  messageID,
		"for_album":   forAlbum,
		"for_comment": forComment,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var messageLink tdlib.MessageLink
	err = json.Unmarshal(result.Raw, &messageLink)
	return &messageLink, err

}
