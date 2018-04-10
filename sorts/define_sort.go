package sorts

type SortInts []int


func (s SortInts)Len() int{
	return len(s)
}

func (s SortInts)Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s SortInts)Less(i, j int)bool {
	return s[i] < s[j]
}
