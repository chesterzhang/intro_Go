package queue

//type Queue []int //相当于将 []int 使用 Queue这个名字来代替
type Queue []interface{}  //Queue 里面可以使用各种类型

func (q *Queue)Push(v int)  {
	*q=append(*q,v)
}

func (q *Queue)Pop()  interface{} {
	head :=(*q)[0]
	*q=(*q)[1:]
	return head
}


func (q *Queue)IsEmpty()  bool{
	return  len(*q)==0
}

