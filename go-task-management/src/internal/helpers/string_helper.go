package helpers

import "strings"

type StringHelper struct{}

func (h *StringHelper) GetSlugString(str string, delimiter *string) string {
	lowerCaseStr := strings.TrimSpace(strings.ToLower(str))
	slicedStr := strings.Split(lowerCaseStr, " ")
	if delimiter != nil {
		return strings.Join(slicedStr, *delimiter)
	}
	return strings.Join(slicedStr, "")
}
