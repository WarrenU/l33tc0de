func compress(chars []byte) int {
	read := 0
	write := 0

	for read < len(chars) {
		c := chars[read]
		count := 0

		for read < len(chars) && chars[read] == c {
			read++
			count++
		}

		chars[write] = c
		write++

		if count > 1 {
			digits := []byte(fmt.Sprintf("%d", count))
			for _, d := range digits {
				chars[write] = d
				write++
			}
		}
	}

	return write
}