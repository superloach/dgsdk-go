package dgsdk

/*
#include "src/c/discord_game_sdk.h"
*/
import "C"

type User C.struct_DiscordUser

func NewUser(
	id UserId,
	username string,
	discriminator string,
	avatar string,
	bot bool,
) *User {
	u := &User{}

	u.SetId(id)
	u.SetUsername(username)
	u.SetDiscriminator(discriminator)
	u.SetAvatar(avatar)
	u.SetBot(bot)

	return u
}

func (u User) Id() UserId {
	return u.id
}

func (u *User) SetId(id UserId) {
	u.id = id
}

func (u User) Username() string {
	username := make([]byte, 0)
	for _, b := range u.username {
		if b == 0 {
			break
		}
		username = append(username, byte(b))
	}
	return string(username)
}

func (u *User) SetUsername(username string) {
	src := []byte(username)
	for i, c := range src {
		u.username[i] = C.char(c)
	}
	u.username[len(src)] = 0
}
