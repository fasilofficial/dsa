// stack.dart

class Stack {
  List<dynamic> _items = [];

  bool get isEmpty => _items.isEmpty;

  void push(dynamic item) {
    _items.add(item);
  }

  dynamic pop() {
    if (isEmpty) {
      throw StateError("Stack is empty");
    }
    return _items.removeLast();
  }

  dynamic peek() {
    if (isEmpty) {
      throw StateError("Stack is empty");
    }
    return _items.last;
  }

  void display() {
    print("Stack: $_items");
  }
}

void main() {
  Stack myStack = Stack();

  myStack.push(1);
  myStack.push(2);
  myStack.push(3);

  print("Original Stack:");
  myStack.display();

  print("\nPop operation result: ${myStack.pop()}");

  print("Stack after Pop operation:");
  myStack.display();
}
