package set

func UniqueStrings(ss []string) []string {
	set := make(map[string]struct{})
	uniqueRes := make([]string, 0)

	// For every el of given slice check wether it was already met as key
	for _, s := range ss {
		
		// If slot for current key was empty fullfil it with struct{}{} and append to slice of unique elems
		if _, ok := set[s]; !ok { 
			set[s] = struct{}{}
			uniqueRes = append(uniqueRes, s)
		}
	}

	return uniqueRes
}
