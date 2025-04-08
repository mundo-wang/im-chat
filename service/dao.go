package service

import "im-chat/dao/query"

var (
	usersQ       = query.Users
	contactsQ    = query.Contacts
	communitiesQ = query.Communities
	messagesQ    = query.Messages
)

func InitDao() {
	usersQ = query.Users
	contactsQ = query.Contacts
	communitiesQ = query.Communities
	messagesQ = query.Messages
}
