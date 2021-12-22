// AUTOGENERATED - DO NOT EDIT

package tdlib

import (
	"encoding/json"
)

// Message Describes a message
type Message struct {
	tdCommon
	ID                        int64                   `json:"id"`                            // Message identifier; unique for the chat to which the message belongs
	Sender                    MessageSender           `json:"sender"`                        // The sender of the message
	ChatID                    int64                   `json:"chat_id"`                       // Chat identifier
	SendingState              MessageSendingState     `json:"sending_state"`                 // Information about the sending state of the message; may be null
	SchedulingState           MessageSchedulingState  `json:"scheduling_state"`              // Information about the scheduling state of the message; may be null
	IsOutgoing                bool                    `json:"is_outgoing"`                   // True, if the message is outgoing
	IsPinned                  bool                    `json:"is_pinned"`                     // True, if the message is pinned
	CanBeEdited               bool                    `json:"can_be_edited"`                 // True, if the message can be edited. For live location and poll messages this fields shows whether editMessageLiveLocation or stopPoll can be used with this message by the application
	CanBeForwarded            bool                    `json:"can_be_forwarded"`              // True, if the message can be forwarded
	CanBeDeletedOnlyForSelf   bool                    `json:"can_be_deleted_only_for_self"`  // True, if the message can be deleted only for the current user while other users will continue to see it
	CanBeDeletedForAllUsers   bool                    `json:"can_be_deleted_for_all_users"`  // True, if the message can be deleted for all users
	CanGetStatistics          bool                    `json:"can_get_statistics"`            // True, if the message statistics are available
	CanGetMessageThread       bool                    `json:"can_get_message_thread"`        // True, if the message thread info is available
	CanGetMediaTimestampLinks bool                    `json:"can_get_media_timestamp_links"` // True, if media timestamp links can be generated for media timestamp entities in the message text, caption or web page description
	HasTimestampedMedia       bool                    `json:"has_timestamped_media"`         // True, if media timestamp entities refers to a media in this message as opposed to a media in the replied message
	IsChannelPost             bool                    `json:"is_channel_post"`               // True, if the message is a channel post. All messages to channels are channel posts, all other messages are not channel posts
	ContainsUnreadMention     bool                    `json:"contains_unread_mention"`       // True, if the message contains an unread mention for the current user
	Date                      int32                   `json:"date"`                          // Point in time (Unix timestamp) when the message was sent
	EditDate                  int32                   `json:"edit_date"`                     // Point in time (Unix timestamp) when the message was last edited
	ForwardInfo               *MessageForwardInfo     `json:"forward_info"`                  // Information about the initial message sender; may be null
	InteractionInfo           *MessageInteractionInfo `json:"interaction_info"`              // Information about interactions with the message; may be null
	ReplyInChatID             int64                   `json:"reply_in_chat_id"`              // If non-zero, the identifier of the chat to which the replied message belongs; Currently, only messages in the Replies chat can have different reply_in_chat_id and chat_id
	ReplyToMessageID          int64                   `json:"reply_to_message_id"`           // If non-zero, the identifier of the message this message is replying to; can be the identifier of a deleted message
	MessageThreadID           int64                   `json:"message_thread_id"`             // If non-zero, the identifier of the message thread the message belongs to; unique within the chat to which the message belongs
	TTL                       int32                   `json:"ttl"`                           // For self-destructing messages, the message's TTL (Time To Live), in seconds; 0 if none. TDLib will send updateDeleteMessages or updateMessageContent once the TTL expires
	TTLExpiresIn              float64                 `json:"ttl_expires_in"`                // Time left before the message expires, in seconds. If the TTL timer isn't started yet, equals to the value of the ttl field
	ViaBotUserID              int64                   `json:"via_bot_user_id"`               // If non-zero, the user identifier of the bot through which this message was sent
	AuthorSignature           string                  `json:"author_signature"`              // For channel posts and anonymous group messages, optional author signature
	MediaAlbumID              JSONInt64               `json:"media_album_id"`                // Unique identifier of an album this message belongs to. Only audios, documents, photos and videos can be grouped together in albums
	RestrictionReason         string                  `json:"restriction_reason"`            // If non-empty, contains a human-readable description of the reason why access to this message must be restricted
	Content                   MessageContent          `json:"content"`                       // Content of the message
	ReplyMarkup               ReplyMarkup             `json:"reply_markup"`                  // Reply markup for the message; may be null
}

// MessageType return the string telegram-type of Message
func (message *Message) MessageType() string {
	return "message"
}

// NewMessage creates a new Message
//
// @param iD Message identifier; unique for the chat to which the message belongs
// @param sender The sender of the message
// @param chatID Chat identifier
// @param sendingState Information about the sending state of the message; may be null
// @param schedulingState Information about the scheduling state of the message; may be null
// @param isOutgoing True, if the message is outgoing
// @param isPinned True, if the message is pinned
// @param canBeEdited True, if the message can be edited. For live location and poll messages this fields shows whether editMessageLiveLocation or stopPoll can be used with this message by the application
// @param canBeForwarded True, if the message can be forwarded
// @param canBeDeletedOnlyForSelf True, if the message can be deleted only for the current user while other users will continue to see it
// @param canBeDeletedForAllUsers True, if the message can be deleted for all users
// @param canGetStatistics True, if the message statistics are available
// @param canGetMessageThread True, if the message thread info is available
// @param canGetMediaTimestampLinks True, if media timestamp links can be generated for media timestamp entities in the message text, caption or web page description
// @param hasTimestampedMedia True, if media timestamp entities refers to a media in this message as opposed to a media in the replied message
// @param isChannelPost True, if the message is a channel post. All messages to channels are channel posts, all other messages are not channel posts
// @param containsUnreadMention True, if the message contains an unread mention for the current user
// @param date Point in time (Unix timestamp) when the message was sent
// @param editDate Point in time (Unix timestamp) when the message was last edited
// @param forwardInfo Information about the initial message sender; may be null
// @param interactionInfo Information about interactions with the message; may be null
// @param replyInChatID If non-zero, the identifier of the chat to which the replied message belongs; Currently, only messages in the Replies chat can have different reply_in_chat_id and chat_id
// @param replyToMessageID If non-zero, the identifier of the message this message is replying to; can be the identifier of a deleted message
// @param messageThreadID If non-zero, the identifier of the message thread the message belongs to; unique within the chat to which the message belongs
// @param tTL For self-destructing messages, the message's TTL (Time To Live), in seconds; 0 if none. TDLib will send updateDeleteMessages or updateMessageContent once the TTL expires
// @param tTLExpiresIn Time left before the message expires, in seconds. If the TTL timer isn't started yet, equals to the value of the ttl field
// @param viaBotUserID If non-zero, the user identifier of the bot through which this message was sent
// @param authorSignature For channel posts and anonymous group messages, optional author signature
// @param mediaAlbumID Unique identifier of an album this message belongs to. Only audios, documents, photos and videos can be grouped together in albums
// @param restrictionReason If non-empty, contains a human-readable description of the reason why access to this message must be restricted
// @param content Content of the message
// @param replyMarkup Reply markup for the message; may be null
func NewMessage(iD int64, sender MessageSender, chatID int64, sendingState MessageSendingState, schedulingState MessageSchedulingState, isOutgoing bool, isPinned bool, canBeEdited bool, canBeForwarded bool, canBeDeletedOnlyForSelf bool, canBeDeletedForAllUsers bool, canGetStatistics bool, canGetMessageThread bool, canGetMediaTimestampLinks bool, hasTimestampedMedia bool, isChannelPost bool, containsUnreadMention bool, date int32, editDate int32, forwardInfo *MessageForwardInfo, interactionInfo *MessageInteractionInfo, replyInChatID int64, replyToMessageID int64, messageThreadID int64, tTL int32, tTLExpiresIn float64, viaBotUserID int64, authorSignature string, mediaAlbumID JSONInt64, restrictionReason string, content MessageContent, replyMarkup ReplyMarkup) *Message {
	messageTemp := Message{
		tdCommon:                  tdCommon{Type: "message"},
		ID:                        iD,
		Sender:                    sender,
		ChatID:                    chatID,
		SendingState:              sendingState,
		SchedulingState:           schedulingState,
		IsOutgoing:                isOutgoing,
		IsPinned:                  isPinned,
		CanBeEdited:               canBeEdited,
		CanBeForwarded:            canBeForwarded,
		CanBeDeletedOnlyForSelf:   canBeDeletedOnlyForSelf,
		CanBeDeletedForAllUsers:   canBeDeletedForAllUsers,
		CanGetStatistics:          canGetStatistics,
		CanGetMessageThread:       canGetMessageThread,
		CanGetMediaTimestampLinks: canGetMediaTimestampLinks,
		HasTimestampedMedia:       hasTimestampedMedia,
		IsChannelPost:             isChannelPost,
		ContainsUnreadMention:     containsUnreadMention,
		Date:                      date,
		EditDate:                  editDate,
		ForwardInfo:               forwardInfo,
		InteractionInfo:           interactionInfo,
		ReplyInChatID:             replyInChatID,
		ReplyToMessageID:          replyToMessageID,
		MessageThreadID:           messageThreadID,
		TTL:                       tTL,
		TTLExpiresIn:              tTLExpiresIn,
		ViaBotUserID:              viaBotUserID,
		AuthorSignature:           authorSignature,
		MediaAlbumID:              mediaAlbumID,
		RestrictionReason:         restrictionReason,
		Content:                   content,
		ReplyMarkup:               replyMarkup,
	}

	return &messageTemp
}

// UnmarshalJSON unmarshal to json
func (message *Message) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		ID                        int64                   `json:"id"`                            // Message identifier; unique for the chat to which the message belongs
		ChatID                    int64                   `json:"chat_id"`                       // Chat identifier
		IsOutgoing                bool                    `json:"is_outgoing"`                   // True, if the message is outgoing
		IsPinned                  bool                    `json:"is_pinned"`                     // True, if the message is pinned
		CanBeEdited               bool                    `json:"can_be_edited"`                 // True, if the message can be edited. For live location and poll messages this fields shows whether editMessageLiveLocation or stopPoll can be used with this message by the application
		CanBeForwarded            bool                    `json:"can_be_forwarded"`              // True, if the message can be forwarded
		CanBeDeletedOnlyForSelf   bool                    `json:"can_be_deleted_only_for_self"`  // True, if the message can be deleted only for the current user while other users will continue to see it
		CanBeDeletedForAllUsers   bool                    `json:"can_be_deleted_for_all_users"`  // True, if the message can be deleted for all users
		CanGetStatistics          bool                    `json:"can_get_statistics"`            // True, if the message statistics are available
		CanGetMessageThread       bool                    `json:"can_get_message_thread"`        // True, if the message thread info is available
		CanGetMediaTimestampLinks bool                    `json:"can_get_media_timestamp_links"` // True, if media timestamp links can be generated for media timestamp entities in the message text, caption or web page description
		HasTimestampedMedia       bool                    `json:"has_timestamped_media"`         // True, if media timestamp entities refers to a media in this message as opposed to a media in the replied message
		IsChannelPost             bool                    `json:"is_channel_post"`               // True, if the message is a channel post. All messages to channels are channel posts, all other messages are not channel posts
		ContainsUnreadMention     bool                    `json:"contains_unread_mention"`       // True, if the message contains an unread mention for the current user
		Date                      int32                   `json:"date"`                          // Point in time (Unix timestamp) when the message was sent
		EditDate                  int32                   `json:"edit_date"`                     // Point in time (Unix timestamp) when the message was last edited
		ForwardInfo               *MessageForwardInfo     `json:"forward_info"`                  // Information about the initial message sender; may be null
		InteractionInfo           *MessageInteractionInfo `json:"interaction_info"`              // Information about interactions with the message; may be null
		ReplyInChatID             int64                   `json:"reply_in_chat_id"`              // If non-zero, the identifier of the chat to which the replied message belongs; Currently, only messages in the Replies chat can have different reply_in_chat_id and chat_id
		ReplyToMessageID          int64                   `json:"reply_to_message_id"`           // If non-zero, the identifier of the message this message is replying to; can be the identifier of a deleted message
		MessageThreadID           int64                   `json:"message_thread_id"`             // If non-zero, the identifier of the message thread the message belongs to; unique within the chat to which the message belongs
		TTL                       int32                   `json:"ttl"`                           // For self-destructing messages, the message's TTL (Time To Live), in seconds; 0 if none. TDLib will send updateDeleteMessages or updateMessageContent once the TTL expires
		TTLExpiresIn              float64                 `json:"ttl_expires_in"`                // Time left before the message expires, in seconds. If the TTL timer isn't started yet, equals to the value of the ttl field
		ViaBotUserID              int64                   `json:"via_bot_user_id"`               // If non-zero, the user identifier of the bot through which this message was sent
		AuthorSignature           string                  `json:"author_signature"`              // For channel posts and anonymous group messages, optional author signature
		MediaAlbumID              JSONInt64               `json:"media_album_id"`                // Unique identifier of an album this message belongs to. Only audios, documents, photos and videos can be grouped together in albums
		RestrictionReason         string                  `json:"restriction_reason"`            // If non-empty, contains a human-readable description of the reason why access to this message must be restricted

	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	message.tdCommon = tempObj.tdCommon
	message.ID = tempObj.ID
	message.ChatID = tempObj.ChatID
	message.IsOutgoing = tempObj.IsOutgoing
	message.IsPinned = tempObj.IsPinned
	message.CanBeEdited = tempObj.CanBeEdited
	message.CanBeForwarded = tempObj.CanBeForwarded
	message.CanBeDeletedOnlyForSelf = tempObj.CanBeDeletedOnlyForSelf
	message.CanBeDeletedForAllUsers = tempObj.CanBeDeletedForAllUsers
	message.CanGetStatistics = tempObj.CanGetStatistics
	message.CanGetMessageThread = tempObj.CanGetMessageThread
	message.CanGetMediaTimestampLinks = tempObj.CanGetMediaTimestampLinks
	message.HasTimestampedMedia = tempObj.HasTimestampedMedia
	message.IsChannelPost = tempObj.IsChannelPost
	message.ContainsUnreadMention = tempObj.ContainsUnreadMention
	message.Date = tempObj.Date
	message.EditDate = tempObj.EditDate
	message.ForwardInfo = tempObj.ForwardInfo
	message.InteractionInfo = tempObj.InteractionInfo
	message.ReplyInChatID = tempObj.ReplyInChatID
	message.ReplyToMessageID = tempObj.ReplyToMessageID
	message.MessageThreadID = tempObj.MessageThreadID
	message.TTL = tempObj.TTL
	message.TTLExpiresIn = tempObj.TTLExpiresIn
	message.ViaBotUserID = tempObj.ViaBotUserID
	message.AuthorSignature = tempObj.AuthorSignature
	message.MediaAlbumID = tempObj.MediaAlbumID
	message.RestrictionReason = tempObj.RestrictionReason

	fieldSender, _ := unmarshalMessageSender(objMap["sender"])
	message.Sender = fieldSender

	fieldSendingState, _ := unmarshalMessageSendingState(objMap["sending_state"])
	message.SendingState = fieldSendingState

	fieldSchedulingState, _ := unmarshalMessageSchedulingState(objMap["scheduling_state"])
	message.SchedulingState = fieldSchedulingState

	fieldContent, _ := unmarshalMessageContent(objMap["content"])
	message.Content = fieldContent

	fieldReplyMarkup, _ := unmarshalReplyMarkup(objMap["reply_markup"])
	message.ReplyMarkup = fieldReplyMarkup

	return nil
}
