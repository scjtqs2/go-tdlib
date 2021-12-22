// AUTOGENERATED - DO NOT EDIT

package client

import (
	"encoding/json"
	"fmt"

	"github.com/Arman92/go-tdlib/v2/tdlib"
)

// GetAllPassportElements Returns all available Telegram Passport elements
// @param password Password of the current user
func (client *Client) GetAllPassportElements(password string) (*tdlib.PassportElements, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type":    "getAllPassportElements",
		"password": password,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var passportElements tdlib.PassportElements
	err = json.Unmarshal(result.Raw, &passportElements)
	return &passportElements, err

}
