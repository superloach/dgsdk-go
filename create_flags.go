/*generated by internal*/
package dgsdk

/*
#include "src/c/discord_game_sdk.h"
*/
import "C"

type CreateFlags C.enum_EDiscordCreateFlags

const (
	CreateFlagsDefault          CreateFlags = C.DiscordCreateFlags_Default
	CreateFlagsNoRequireDiscord             = C.DiscordCreateFlags_NoRequireDiscord
)