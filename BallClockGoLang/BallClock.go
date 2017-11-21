package main

import (
	"fmt"
	"strconv"
	"strings"
	"bufio"
	"os"
)

// Stacks and Queues  https://gist.github.com/moraes/2141121
type Node struct {
	Value int
}

func (n *Node) String() string {
	return fmt.Sprint(n.Value)
}

// NewStack returns a new stack.
func NewStack() *Stack {
	return &Stack{}
}

// Stack is a basic LIFO stack that resizes as needed.
type Stack struct {
	nodes []*Node
	count int
}

// Push adds a node to the stack.
func (s *Stack) Push(n *Node) {
	s.nodes = append(s.nodes[:s.count], n)
	s.count++
}

// Pop removes and returns a node from the stack in last to first order.
func (s *Stack) Pop() *Node {
	if s.count == 0 {
		return nil
	}
	s.count--
	return s.nodes[s.count]
}

// NewQueue returns a new queue with the given initial size.
func NewQueue(size int) *Queue {
	return &Queue{
		nodes: make([]*Node, size),
		size:  size,
	}
}

// Queue is a basic FIFO queue based on a circular list that resizes as needed.
type Queue struct {
	nodes []*Node
	size  int
	head  int
	tail  int
	count int
}

// Push adds a node to the queue.
func (q *Queue) Push(n *Node) {
	if q.head == q.tail && q.count > 0 {
		nodes := make([]*Node, len(q.nodes)+q.size)
		copy(nodes, q.nodes[q.head:])
		copy(nodes[len(q.nodes)-q.head:], q.nodes[:q.head])
		q.head = 0
		q.tail = len(q.nodes)
		q.nodes = nodes
	}
	q.nodes[q.tail] = n
	q.tail = (q.tail + 1) % len(q.nodes)
	q.count++
}

// Pop removes and returns a node from the queue in first to last order.
func (q *Queue) Pop() *Node {
	if q.count == 0 {
		return nil
	}
	node := q.nodes[q.head]
	q.head = (q.head + 1) % len(q.nodes)
	q.count--
	return node
}

var twelveCount = 0
var minuteCount = 0
var dayCount = 0
var numBalls = -1
var numberMinutesToRun = -1
var originalOrder = ""
var originalId = ""
var minuteStack = NewStack()
var fiveStack = NewStack()
var hourStack = NewStack()
var ballQueue = NewQueue(1)



func main() {

	input:=""
	print("Enter number of balls: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input = scanner.Text()

	if _, err := fmt.Sscan(input, &numBalls); err == nil {
		//fmt.Println("sucess")
	}else{
		fmt.Println("invalid input")
		os.Exit(3)
	}

	print("Do you want to enter minutes? (y/n)")
	scanner2 := bufio.NewScanner(os.Stdin)
	scanner2.Scan()
	input = scanner2.Text()


	if input == "y" || input == "Y"{
		print("Enter Minutes: ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input = scanner.Text()

		if _, err := fmt.Sscan(input, &numberMinutesToRun); err == nil {
			//fmt.Println("sucess")
		}else{
			fmt.Println("invalid input")
			os.Exit(3)
		}

	}


	hourStack.Push(&Node{0})
	loadQueue(numBalls)
	originalOrder = getCurrentOrder()

	if(numberMinutesToRun == -1){

		loadStacks()

		for i:=0; i<1;{

			executeMinute()
			if originalOrder == getCurrentOrder(){
				break
			}

		}

		fmt.Println(strconv.Itoa(numBalls) + " balls cycle after "+strconv.Itoa(dayCount)+ " days")
	}else{
		executeMinutes(numberMinutesToRun);
		fmt.Println(printJson())
	}


}//end main



func loadQueue(numBalls int ){
	for i :=1; i<= numBalls; i++{
		ballQueue.Push(&Node{i})
	}
}

func loadStacks() {

	for i:=1; i<12; i++{
		var temp = ballQueue.Pop()
		hourStack.Push(temp)
	}


	for i:=0; i<11; i++{
		var temp = ballQueue.Pop()
		fiveStack.Push(temp)
	}

	for i:=0; i<4; i++{
		var temp = ballQueue.Pop()
		minuteStack.Push(temp)
	}

}//endLoadStacks

func executeMinute() {


	if len(minuteStack.nodes) != 4 {
		minuteStack.Push(ballQueue.Pop());//if there is room in the minuteStack, add a ball from the queue.
	}else {

		for i:=0; i<4; i++ {

			ballQueue.Push(minuteStack.Pop())
		}//pop all off the minute stack, add to ballQueue

		minuteStack = NewStack()

		if len(fiveStack.nodes) != 11 {
			fiveStack.Push(ballQueue.Pop())
		}else {

			for j:=0; j<11; j++{
				ballQueue.Push(fiveStack.Pop())
			}
			fiveStack=NewStack()
			if len(hourStack.nodes) != 12 {
				hourStack.Push(ballQueue.Pop());
			}else {

				for k:=0; k<11; k++{
					ballQueue.Push(hourStack.Pop());
				}
				hourStack = NewStack()
				hourStack.Push(&Node{0})
				ballQueue.Push(ballQueue.Pop());

				twelveCount++

				if twelveCount == 2{
					dayCount++
					//fmt.Println("Days: "+ strconv.Itoa(dayCount))
					twelveCount =0
				}

			} //end hourStack else

		} //end fiveStack else

	} //end if/else if

	//fmt.Println(getTime())

}//end executeMinute

func executeMinutes(minutes int) {

	for x:=0; x<minutes; x++ {

		if len(minuteStack.nodes) != 4 {
			minuteStack.Push(ballQueue.Pop());//if there is room in the minuteStack, add a ball from the queue.
		}else {

			for i:=0; i<4; i++ {

				ballQueue.Push(minuteStack.Pop())
			}//pop all off the minute stack, add to ballQueue

			minuteStack = NewStack()

			if len(fiveStack.nodes) != 11 {
				fiveStack.Push(ballQueue.Pop())//if there is room in the Fives Stack, add ball from the queue.
			}else {

				for j:=0; j<11; j++{
					ballQueue.Push(fiveStack.Pop())
				}//pop all off the fiveStack, and add to ballQueue
				fiveStack=NewStack()
				if len(hourStack.nodes) != 12 {
					hourStack.Push(ballQueue.Pop());
				}else {

					for k:=0; k<11; k++{
						ballQueue.Push(hourStack.Pop());
					}
					hourStack = NewStack()
					hourStack.Push(&Node{0})
					ballQueue.Push(ballQueue.Pop());

					twelveCount++

					if twelveCount == 2{
						dayCount++
						twelveCount =0
					}

				} //end hourStack else

			} //end fiveStack else

		} //end if/else if

		//fmt.Println(getTime())
	}

}//end executeMinutes

func getTime() string {

	hour := strconv.Itoa(len(hourStack.nodes))
	minutes := ""


	if len(fiveStack.nodes) ==1 || len(fiveStack.nodes) ==0{
		minutes = "0"+ strconv.Itoa(len(minuteStack.nodes))
	}else{
		minutes = strconv.Itoa(len(fiveStack.nodes)*5 + len(minuteStack.nodes));
	}


return hour + ":" + minutes


}//end getTime()

func getCurrentOrder() string{
	currentOrder:=""

	for i:=1; i<len(hourStack.nodes);i++{
		node := hourStack.nodes[i]
		currentOrder+=node.String()
	}

	if len(fiveStack.nodes) != 0 {
		for i:=0; i<len(fiveStack.nodes); i++ {
			node :=fiveStack.nodes[i]
			currentOrder+=node.String()
		}
	}

	if len(minuteStack.nodes) != 0 {
		for i:=0; i<len(minuteStack.nodes);i++{
			node := minuteStack.nodes[i]
			currentOrder+=node.String()
		}
	}

	count:=0
	for i:=ballQueue.head; i< len(ballQueue.nodes); i++{
		node := ballQueue.nodes[i]
		if(node != nil) {

			currentOrder += node.String()
			count++
		}
	}

	//since this is a circular queue, we may not cature all the needed values. so we will get the values at the beginning by doing the following:
	//num := len(minuteStack.nodes)+len(fiveStack.nodes)+len(hourStack.nodes)+count-1

	////second loop from the beginning
	//for i:=0; i < num; i++{
	//	node := ballQueue.nodes[i]
	//	if(node != nil) {
	//
	//		currentOrder += node.String() +","
	//	}
	//}

	return currentOrder
}

func printJson() string{

	hours :=""
	for i:=1; i< len(hourStack.nodes); i++{
		node := hourStack.nodes[i]
		if i == len(hourStack.nodes)-1{
			hours += node.String()
		}else{
			hours += node.String() +","
		}

	}

	fives :=""
	for i:=0; i< len(fiveStack.nodes); i++{
		node := fiveStack.nodes[i]
		if i == len(fiveStack.nodes)-1{
			fives += node.String()
		}else{
			fives += node.String() +","
		}

	}

	min:=""
	for i:=1; i< len(minuteStack.nodes); i++{
		node := minuteStack.nodes[i]
		if i == len(minuteStack.nodes)-1{
			min+= node.String()
		}else{
			min+= node.String() +","
		}

	}

	count:=0

	queue:=""
	for i:=ballQueue.head; i< len(ballQueue.nodes); i++{
		node := ballQueue.nodes[i]
		if(node != nil) {

			queue += node.String() +","
			count++
		}
	}

	//since this is a circular queue, we may not cature all the needed values. so we will get the values at the beginning by doing the following:
	num := len(minuteStack.nodes)+len(fiveStack.nodes)+len(hourStack.nodes)+count-1

	//second loop from the beginning
	for i:=0; i < num; i++{
		node := ballQueue.nodes[i]
		if(node != nil) {

			queue += node.String() +","
		}
	}

	if queue != "" {
		queue = strings.TrimSuffix(queue, ",")
	}



	json:="{\"Min\":["+min+"],\"FiveMin\":["+fives+"],\"Hour\":["+hours+"],\"Main\"["+queue+"]}"
	//json = json.replaceAll(" ", "");
	return json


}
