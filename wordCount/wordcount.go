package wordCount

func WordCount(str string) uint32 {
	var i = 0
	var words uint32 = 0

	for ; i < len(str); i++ {
		if str[i] == ' ' || str[i] == '\t' || str[i] == '\n' {
			continue
		}
		for ; i < len(str); i++ {
			if str[i] == ' '  || str[i] == '\t' || str[i] == '\n'{
				break
			}
		}
		words++
	}

	return words
}
