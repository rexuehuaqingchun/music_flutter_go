package constant

const (
	DefaultPageCount = 15
)

var (
	UserRoles = map[string]string{
		"NORMAL_USER":         "NORMAL_USER",
		"DQ_SINGER":           "DQ_SINGER",
		"DQ_OFFICIAL_ACCOUNT": "DQ_OFFICIAL_ACCOUNT",
	}
)

func DetectionUserRoles(userRoles string) bool {
	_, ok := UserRoles[userRoles]
	return ok
}
