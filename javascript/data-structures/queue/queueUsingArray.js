class Queue{
  constructor(){
    this.items = [];
  }

  enqueue(element){
    this.items.push(element);
  }

  dequeue(){
    if (!this.isEmpty()) {
    return this.items.shift();
    }
    return 'Queue is empty'
  }

  peek(){
    if (!this.isEmpty()) {
    return this.items[0];
    }
    return "Queue is empty"
  }

  isEmpty(){
    return this.items.length == 0;
  }

  print(){
    console.log(this.items);
  }
}

//Create new Queue
const myQueue = new Queue();

// Check if the queue is empty
console.log("Is the queue empty?", myQueue.isEmpty());

// Peek the first element
console.log("First element: ", myQueue.peek());

// Push elements onto the queue
myQueue.enqueue(5);
myQueue.enqueue(10);
myQueue.enqueue(15);

// Display the queue's content
console.log("Queue after push operations:");
myQueue.print();

// Peek the first element
console.log("First element: ", myQueue.peek());

// Removed an element
const removedElement = myQueue.dequeue();
console.log("Removed element:", removedElement);

// Check if the queue is empty
console.log("Is the queue empty?", myQueue.isEmpty());

// Display the queue's content after pop
console.log("Queue after dequeue operation:");
myQueue.print();