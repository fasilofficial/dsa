class MyArray {
  late List<int> _elements;

  // Constructor
  MyArray(int size) {
    _elements = List<int>.filled(size, 0);
  }

  // Get the length of the array
  int get length => _elements.length;

  // Access element at index
  int getElementAt(int index) {
    if (index < 0 || index >= _elements.length) {
      throw RangeError('Index out of bounds');
    }
    return _elements[index];
  }

  // Insert element at the start of the array
  void insertAtStart(int value) {
    _elements = [value, ..._elements];
  }

  // Insert element at the end of the array
  void insertAtEnd(int value) {
    _elements.add(value);
  }

  // Insert element at a specific position
  void insertAtPosition(int index, int value) {
    if (index < 0 || index > _elements.length) {
      throw RangeError('Index out of bounds');
    }
    _elements.insert(index, value);
  }

  // Delete element at a specific position
  void deleteAtPosition(int index) {
    if (index < 0 || index >= _elements.length) {
      throw RangeError('Index out of bounds');
    }
    _elements.removeAt(index);
  }

  // Print the elements of the array
  void printArray() {
    print(_elements);
  }
}

void main() {
  // Example usage of array
  MyArray myArray = MyArray(5);

  myArray.insertAtStart(0);
  myArray.insertAtEnd(2);
  myArray.insertAtPosition(1, 1);

  print('Array after insertions:');
  myArray.printArray();

  myArray.deleteAtPosition(1);

  print('Array after deletion:');
  myArray.printArray();
}
