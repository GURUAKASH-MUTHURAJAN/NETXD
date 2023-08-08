package main

import("fmt")
type Student struct {
    Name      string
    RollNo    int
}
type Queue struct{
	items []interface{}
}
func (q *Queue) isempty() bool{
	return len(q.items)==0
}
func (q *Queue) Enqueue(val interface{}){
	q.items=append(q.items,val)
}
func (q *Queue) Dequeue() interface{}{
	if len(q.items)==0{
		return nil
	}
	value:=q.items[0]
	q.items=q.items[1:]
	return value
}
func main(){
	queue:=&Queue{}
	fmt.Println("Is Queue Empty",queue.isempty())
	student1 := Student{Name: "Alice", RollNo: 101}
    student2 := Student{Name: "Bob", RollNo: 102}
    student3 := Student{Name: "Charlie", RollNo: 103}

    queue.Enqueue(student1)
    queue.Enqueue(student2)
    queue.Enqueue(student3)
	fmt.Println("Is Queue Empty",queue.isempty())
	fmt.Println(*queue)
	fmt.Println("Dequeue: ",queue.Dequeue())
	fmt.Println("Dequeue: ",queue.Dequeue())
	fmt.Println("Dequeue: ",queue.Dequeue())
	fmt.Println("Dequeue: ",queue.Dequeue())
	fmt.Println("Dequeue: ",queue.Dequeue())
	fmt.Println("Is Queue Empty",queue.isempty())
}
