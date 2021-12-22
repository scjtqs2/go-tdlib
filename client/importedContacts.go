// AUTOGENERATED - DO NOT EDIT

package client

import (
	"encoding/json"
	"fmt"

	"github.com/Arman92/go-tdlib/v2/tdlib"
)

// ImportContacts Adds new contacts or edits existing contacts by their phone numbers; contacts' user identifiers are ignored
// @param contacts The list of contacts to import or edit; contacts' vCard are ignored and are not imported
func (client *Client) ImportContacts(contacts []tdlib.Contact) (*tdlib.ImportedContacts, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type":    "importContacts",
		"contacts": contacts,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var importedContacts tdlib.ImportedContacts
	err = json.Unmarshal(result.Raw, &importedContacts)
	return &importedContacts, err

}

// ChangeImportedContacts Changes imported contacts using the list of current user contacts saved on the device. Imports newly added contacts and, if at least the file database is enabled, deletes recently deleted contacts. Query result depends on the result of the previous query, so only one query is possible at the same time
// @param contacts The new list of contacts, contact's vCard are ignored and are not imported
func (client *Client) ChangeImportedContacts(contacts []tdlib.Contact) (*tdlib.ImportedContacts, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type":    "changeImportedContacts",
		"contacts": contacts,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var importedContacts tdlib.ImportedContacts
	err = json.Unmarshal(result.Raw, &importedContacts)
	return &importedContacts, err

}
