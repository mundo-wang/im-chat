package service

import "im-chat/dao/query"

var (
	usersQ       = query.Users
	contactsQ    = query.Contacts
	communitiesQ = query.Communities
)

func InitDao() {
	usersQ = query.Users
	contactsQ = query.Contacts
	communitiesQ = query.Communities
}
