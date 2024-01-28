class Queue{
  constructor(){
    this.items = {};
    this.front = 0;
    this.rear = 0;
  }

  enqueue(element){
    this.items[this.rear] = element;
    this.rear++
  }

  dequeue(){
    const item = this.items[this.front];
    delete this.items[this.front];
    this.front++
    return item;
  }

  isEmpty(){
    return this.rear - this.front == 0;
  }

  peek(){
    return this.items[this.front];
  }

  size(){
    return this.rear - this.front;
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

//size of the queue
console.log('Size of queue: ', myQueue.size());

// Push elements onto the queue
myQueue.enqueue(5);
myQueue.enqueue(10);
myQueue.enqueue(15);

//size of the queue
console.log('Size of queue: ', myQueue.size());

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