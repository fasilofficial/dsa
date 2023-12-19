// queue.dart

class Queue {
  List<dynamic> _items = [];

  bool get isEmpty => _items.isEmpty;

  void enqueue(dynamic item) {
    _items.add(item);
  }

  dynamic dequeue() {
    if (isEmpty) {
      throw StateError("Queue is empty");
    }
    return _items.removeAt(0);
  }

  dynamic front() {
    if (isEmpty) {
      throw StateError("Queue is empty");
    }
    return _items.first;
  }

  void display() {
    print("Queue: $_items");
  }
}

void main() {
  Queue myQueue = Queue();

  myQueue.enqueue(1);
  myQueue.enqueue(2);
  myQueue.enqueue(3);

  print("Original Queue:");
  myQueue.display();

  print("\nDequeue operation result: ${myQueue.dequeue()}");

  print("Queue after Dequeue operation:");
  myQueue.display();
}
