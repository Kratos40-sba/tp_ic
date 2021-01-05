package calculs

func Counter(t []string) map[string]int {
	c := map[string]int{}
	for _, v := range t {
		c[v] += 1
	}
	return c
}
