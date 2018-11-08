package resize

func unique(in []string) []string {
	index := map[string]struct{}{}
	out := []string{}
	for _, str := range in {
		if _, ok := index[str]; ok {
			continue
		}
		index[str] = struct{}{}
		out = append(out, str)
	}
	return out
}
