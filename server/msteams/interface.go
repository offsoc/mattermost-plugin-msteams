package msteams

import (
	"context"
	"io"
)

type Client interface {
	Connect() error
	SendMessage(teamID, channelID, parentID, message string) (string, error)
	SendMessageWithAttachments(teamID, channelID, parentID, message string, attachments []*Attachment) (string, error)
	UploadFile(teamID, channelID, filename string, filesize int, mimeType string, data io.Reader) (*Attachment, error)
	UpdateMessage(teamID, channelID, parentID, msgID, message string) error
	DeleteMessage(teamID, channelID, parentID, msgID string) error
	SubscribeToChannel(teamID, channelID, notificationURL, webhookSecret string) (string, error)
	Subscribe(notificationURL, webhookSecret string) (string, error)
	SubscribeToChats(notificationURL, webhookSecret string) (string, error)
	RefreshSubscriptionPeriodically(ctx context.Context, subscriptionID string) error
	ClearSubscription(subscriptionID string) error
	ClearSubscriptions() error
	GetTeam(teamID string) (*Team, error)
	GetChannel(teamID, channelID string) (*Channel, error)
	GetChat(chatID string) (*Chat, error)
	GetChatMessage(chatID, messageID string) (*Message, error)
	GetMessage(teamID, channelID, messageID string) (*Message, error)
	GetReply(teamID, channelID, messageID, replyID string) (*Message, error)
	GetUserAvatar(userID string) ([]byte, error)
	GetMyID() (string, error)
	GetFileURL(weburl string) (string, error)
	GetCodeSnippet(url string) (string, error)
	BotID() string
}
