// AUTOGENERATED - DO NOT EDIT

package client

import (
	"encoding/json"
	"fmt"

	"github.com/Arman92/go-tdlib/v2/tdlib"
)

// SearchChatMembers Searches for a specified query in the first name, last name and username of the members of a specified chat. Requires administrator rights in channels
// @param chatID Chat identifier
// @param query Query to search for
// @param limit The maximum number of users to be returned
// @param filter The type of users to return. By default, chatMembersFilterMembers
func (client *Client) SearchChatMembers(chatID int64, query string, limit int32, filter tdlib.ChatMembersFilter) (*tdlib.ChatMembers, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type":   "searchChatMembers",
		"chat_id": chatID,
		"query":   query,
		"limit":   limit,
		"filter":  filter,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var chatMembers tdlib.ChatMembers
	err = json.Unmarshal(result.Raw, &chatMembers)
	return &chatMembers, err

}

// GetSupergroupMembers Returns information about members or banned users in a supergroup or channel. Can be used only if SupergroupFullInfo.can_get_members == true; additionally, administrator privileges may be required for some filters
// @param supergroupID Identifier of the supergroup or channel
// @param filter The type of users to return. By default, supergroupMembersFilterRecent
// @param offset Number of users to skip
// @param limit The maximum number of users be returned; up to 200
func (client *Client) GetSupergroupMembers(supergroupID int32, filter tdlib.SupergroupMembersFilter, offset int32, limit int32) (*tdlib.ChatMembers, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type":         "getSupergroupMembers",
		"supergroup_id": supergroupID,
		"filter":        filter,
		"offset":        offset,
		"limit":         limit,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var chatMembers tdlib.ChatMembers
	err = json.Unmarshal(result.Raw, &chatMembers)
	return &chatMembers, err

}
