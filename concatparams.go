package piscine

func ConcatParams(args []string) string {
	str := ""
	for i, arg := range args {
		if i > 0 {
			str += "\n"
		}
		str += arg
	}
	return str
}
