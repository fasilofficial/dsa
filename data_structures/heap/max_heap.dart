// max_heap.dart

class MaxHeap {
  List<int> _heap;

  MaxHeap() : _heap = [];

  void insert(int value) {
    _heap.add(value);
    _heapifyUp();
  }

  void _heapifyUp() {
    int index = _heap.length - 1;

    while (index > 0) {
      int parentIndex = (index - 1) ~/ 2;

      if (_heap[index] > _heap[parentIndex]) {
        _swap(index, parentIndex);
        index = parentIndex;
      } else {
        break;
      }
    }
  }

  int extractMax() {
    if (_heap.isEmpty) {
      throw StateError("Heap is empty");
    }

    int maxValue = _heap[0];

    if (_heap.length == 1) {
      _heap.removeAt(0);
    } else {
      _heap[0] = _heap.removeLast();
      _heapifyDown();
    }

    return maxValue;
  }

  void _heapifyDown() {
    int index = 0;

    while (true) {
      int leftChildIndex = 2 * index + 1;
      int rightChildIndex = 2 * index + 2;
      int largestChildIndex = index;

      if (leftChildIndex < _heap.length &&
          _heap[leftChildIndex] > _heap[largestChildIndex]) {
        largestChildIndex = leftChildIndex;
      }

      if (rightChildIndex < _heap.length &&
          _heap[rightChildIndex] > _heap[largestChildIndex]) {
        largestChildIndex = rightChildIndex;
      }

      if (largestChildIndex != index) {
        _swap(index, largestChildIndex);
        index = largestChildIndex;
      } else {
        break;
      }
    }
  }

  void _swap(int i, int j) {
    int temp = _heap[i];
    _heap[i] = _heap[j];
    _heap[j] = temp;
  }

  void display() {
    print("Max Heap: $_heap");
  }
}

void main() {
  MaxHeap maxHeap = MaxHeap();

  maxHeap.insert(3);
  maxHeap.insert(1);
  maxHeap.insert(4);
  maxHeap.insert(2);

  print("Original Max Heap:");
  maxHeap.display();

  print("\nExtract Max Value: ${maxHeap.extractMax()}");

  print("Max Heap after Extraction:");
  maxHeap.display();
}
