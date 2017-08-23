package convert

import (
	"fmt"
	"treasure/helper"
)

func ToFullAvatarUrl(avatar string, domain, defaultAvatar string) string {
	if helper.IsUrl(avatar) {
		return avatar
	}
	if avatar == "" || avatar == "#" {
		return fmt.Sprint(domain, defaultAvatar)
	}
	return fmt.Sprint(domain, avatar)
}
