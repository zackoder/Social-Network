package models

import "strconv"

func UpdateProfile(id int) string {
	privacy, _ := IsPrivateProfile(strconv.Itoa(id))
	updateprofile := "UPDATE users SET privacy = ? WHERE id = ?"
	if !privacy {
		Db.Exec(updateprofile, "privet", id)
		return "privet"
	} else if privacy {
		Db.Exec(updateprofile, "public", id)
		return "public"
	}
	return ""
}
