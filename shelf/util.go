package shelf

import "strings"

//a simple utilities file to help me do some string manipulations

//titleToFile will convert any title with spaces to a file friendly format
func titleToFile(title string) string {
    t := strings.TrimSpace(title)
    t = strings.ToLower(t)
    t = strings.ReplaceAll(t, " ", "_")

    return t
}