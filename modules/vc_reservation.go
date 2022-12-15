package modules

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

// we want to allow people to react to a message, and reserve a VC
// step 1: react to message
// step 2: ask how many other people they want to talk with
// step 3: unlock VC, give them link to join
// (all using ephemeral replies)

// eventually: add dynamic voice channel addition/removal using DB
// for now: use static array of VC IDs

// modifyAccess
func _(session *discordgo.Session, guildID string, channelID string, lock bool) error {
	// basically, deny/allow @everyone the permission to connect to the VC
	// requires that the bot has the "Connect" permission allowed for itself, otherwise it cannot modify permission

	var allowPermissionInt int64
	var denyPermissionInt int64

	if lock {
		// permission int for "Connect" to VC
		denyPermissionInt = 1048576
	} else {
		allowPermissionInt = 1048576
	}

	// we're using guild ID as target ID to specify @everyone
	err := session.ChannelPermissionSet(
		channelID, guildID, discordgo.PermissionOverwriteTypeRole, allowPermissionInt,
		denyPermissionInt,
	)

	if err != nil {
		return fmt.Errorf("[VC ACCESS] unable to set permissions on channel: %w", err)
	}

	return nil
}
