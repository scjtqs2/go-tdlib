// AUTOGENERATED - DO NOT EDIT

package tdlib

// SponsoredMessages Contains a list of sponsored messages
type SponsoredMessages struct {
	tdCommon
	Messages []SponsoredMessage `json:"messages"` // List of sponsored messages
}

// MessageType return the string telegram-type of SponsoredMessages
func (sponsoredMessages *SponsoredMessages) MessageType() string {
	return "sponsoredMessages"
}

// NewSponsoredMessages creates a new SponsoredMessages
//
// @param messages List of sponsored messages
func NewSponsoredMessages(messages []SponsoredMessage) *SponsoredMessages {
	sponsoredMessagesTemp := SponsoredMessages{
		tdCommon: tdCommon{Type: "sponsoredMessages"},
		Messages: messages,
	}

	return &sponsoredMessagesTemp
}
