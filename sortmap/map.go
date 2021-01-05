package sortmap

type Map struct {
	Name  string
	Value int
}
type List []Map

func (l List) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}
func (l List) Len() int {
	return len(l)
}
func (l List) Less(i, j int) bool {
	if l[i].Value == l[j].Value {
		return l[i].Name < l[j].Name
	} else {
		return l[i].Value > l[j].Value
	}
}
