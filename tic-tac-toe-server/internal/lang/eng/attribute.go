package eng

var attribute = map[string]string{
	"user_id":     "User",
	"category_id": "Category",
	"platform_id": "Platform",
	"passowrd":    "Password",
	"email":       "E-mail",
	"name":        "Name",
	"firstname":   "Firstname",
	"lastname":    "Lastname",
	"patronymic":  "Patronymic",
	"text":        "Text",
	"is_private":  "Private",
}

func GetAttribute(field string) string {
	return attribute[field]
}
