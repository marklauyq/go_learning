package shelf

import "strings"

//a simple utilities file to help me do some string manipulations

//titleToFile will convert any title with spaces to a file friendly format
func titleToFile(title string) string {
    t := strings.TrimSpace(title)
    t = strings.ToLower(t)
    t = strings.ReplaceAll(t, " ", "_")
    t = t + ".json"

    return t
}

func titleToFileWithExtensions(title string, ext string) string {
    t := strings.TrimSpace(title)
    t = strings.ToLower(t)
    t = strings.ReplaceAll(t, " ", "_")
    t = t + ext

    return t
}

func cleanString(str string) string {
    t := strings.ReplaceAll(str, "TLN:", "tln-")
    t = strings.ReplaceAll(t, "\n", "<br /><br />")
    t = strings.ReplaceAll(t, "<", "(")
    t = strings.ReplaceAll(t, ">", ")")

    return t
}