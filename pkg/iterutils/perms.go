package iterutils

// ForAllPerms execute cb for on all permutations of input.
func ForAllPerms(input []string, cb func([]string)) {
	s := make([]string, len(input))
	copy(s, input)

	var p func(int)
	p = func(i int) {
		if i > len(s) {
			cb(s)
			return
		}

		p(i + 1)
		for j := i + 1; j < len(s); j++ {
			s[i], s[j] = s[j], s[i]
			p(i + 1)
			s[i], s[j] = s[j], s[i]
		}
	}
	p(0)
}
