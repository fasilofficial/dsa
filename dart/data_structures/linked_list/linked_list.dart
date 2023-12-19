class Node {
  dynamic data;
  Node? next;

  Node(this.data);
}

class SinglyLinkedList {
  Node? head;

  // Insertion at the beginning of the list
  void insertAtBeginning(dynamic data) {
    Node newNode = Node(data);
    newNode.next = head;
    head = newNode;
  }

  // Insertion at the end of the list
  void insertAtEnd(dynamic data) {
    Node newNode = Node(data);
    if (head == null) {
      head = newNode;
    } else {
      Node current = head!;
      while (current.next != null) {
        current = current.next!;
      }
      current.next = newNode;
    }
  }

  // Insertion at a specific position
  void insertAtPosition(dynamic data, int position) {
    Node newNode = Node(data);
    if (position < 0) position = 0;

    if (head == null || position == 0) {
      newNode.next = head;
      head = newNode;
    } else {
      Node current = head!;
      for (int i = 0; i < position - 1 && current.next != null; i++) {
        current = current.next!;
      }
      newNode.next = current.next;
      current.next = newNode;
    }
  }

  // Deletion from the beginning of the list
  void deleteFromBeginning() {
    if (head != null) {
      head = head!.next;
    }
  }

  // Deletion from the end of the list
  void deleteFromEnd() {
    if (head == null) return;

    if (head!.next == null) {
      head = null;
    } else {
      Node current = head!;
      while (current.next!.next != null) {
        current = current.next!;
      }
      current.next = null;
    }
  }

  // Deletion from a specific position
  void deleteFromPosition(int position) {
    if (head == null || position < 0) return;

    if (position == 0) {
      head = head!.next;
    } else {
      Node current = head!;
      for (int i = 0; i < position - 1 && current.next != null; i++) {
        current = current.next!;
      }
      if (current.next != null) {
        current.next = current.next!.next;
      }
    }
  }

  // Display the elements of the list
  void display() {
    Node? current = head;
    while (current != null) {
      print(current.data);
      current = current.next;
    }
  }
}

void main() {
  SinglyLinkedList myList = SinglyLinkedList();

  myList.insertAtEnd(1);
  myList.insertAtEnd(2);
  myList.insertAtEnd(3);

  print("Original List:");
  myList.display();

  myList.insertAtBeginning(0);
  myList.insertAtEnd(4);
  myList.insertAtPosition(2.5, 3);

  print("\nList after Insertions:");
  myList.display();

  myList.deleteFromBeginning();
  myList.deleteFromEnd();
  myList.deleteFromPosition(2);

  print("\nList after Deletions:");
  myList.display();
}
