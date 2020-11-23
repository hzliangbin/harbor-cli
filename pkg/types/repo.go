package types

type Repo struct {
	Name       string
	TagsNum    int
}

type Repos []Repo

func (r Repos)Len() int {
	return len(r)
}

//小根堆
func (r Repos) Less(i,j int) bool {
	return r[i].TagsNum < r[j].TagsNum
}

func (r *Repos) Swap(i,j int) {
	(*r)[i], (*r)[j] = (*r)[j], (*r)[i]
}

func (r *Repos) Push(x interface{}) {
	*r = append(*r, x.(Repo))
}

func (r *Repos) Pop() interface{} {
	res := (*r)[len(*r)-1]
	*r = (*r)[:len(*r)-1]
	return res
}
