package models

import "strconv"

func UpdateProfile(id int) string {
	privacy, _ := GetProfilePrivecy(strconv.Itoa(id))
	updateprofile := "UPDATE users SET privacy = ? WHERE id = ?"
	if privacy == "public" {
		Db.Exec(updateprofile, "privet", id)
		return "privet"
	} else if privacy == "privet" {
		Db.Exec(updateprofile, "public", id)
		return "public"
	}
	return ""
}
