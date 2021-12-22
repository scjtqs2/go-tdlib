// AUTOGENERATED - DO NOT EDIT

package client

import (
	"encoding/json"

	"github.com/Arman92/go-tdlib/v2/tdlib"
)

// RegisterDevice Registers the currently used device for receiving push notifications. Returns a globally unique identifier of the push notification subscription
// @param deviceToken Device token
// @param otherUserIDs List of user identifiers of other users currently using the application
func (client *Client) RegisterDevice(deviceToken tdlib.DeviceToken, otherUserIDs []int64) (*tdlib.PushReceiverID, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type":          "registerDevice",
		"device_token":   deviceToken,
		"other_user_ids": otherUserIDs,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, tdlib.RequestError{Code: int(result.Data["code"].(float64)), Message: result.Data["message"].(string)}
	}

	var pushReceiverID tdlib.PushReceiverID
	err = json.Unmarshal(result.Raw, &pushReceiverID)
	return &pushReceiverID, err

}

// GetPushReceiverID Returns a globally unique push notification subscription identifier for identification of an account, which has received a push notification. Can be called synchronously
// @param payload JSON-encoded push notification payload
func (client *Client) GetPushReceiverID(payload string) (*tdlib.PushReceiverID, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type":   "getPushReceiverId",
		"payload": payload,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, tdlib.RequestError{Code: int(result.Data["code"].(float64)), Message: result.Data["message"].(string)}
	}

	var pushReceiverID tdlib.PushReceiverID
	err = json.Unmarshal(result.Raw, &pushReceiverID)
	return &pushReceiverID, err

}
