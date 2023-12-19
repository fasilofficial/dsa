// min_heap.dart

class MinHeap {
  List<int> _heap;

  MinHeap() : _heap = [];

  void insert(int value) {
    _heap.add(value);
    _heapifyUp();
  }

  void _heapifyUp() {
    int index = _heap.length - 1;

    while (index > 0) {
      int parentIndex = (index - 1) ~/ 2;

      if (_heap[index] < _heap[parentIndex]) {
        _swap(index, parentIndex);
        index = parentIndex;
      } else {
        break;
      }
    }
  }

  int extractMin() {
    if (_heap.isEmpty) {
      throw StateError("Heap is empty");
    }

    int minValue = _heap[0];

    if (_heap.length == 1) {
      _heap.removeAt(0);
    } else {
      _heap[0] = _heap.removeLast();
      _heapifyDown();
    }

    return minValue;
  }

  void _heapifyDown() {
    int index = 0;

    while (true) {
      int leftChildIndex = 2 * index + 1;
      int rightChildIndex = 2 * index + 2;
      int smallestChildIndex = index;

      if (leftChildIndex < _heap.length &&
          _heap[leftChildIndex] < _heap[smallestChildIndex]) {
        smallestChildIndex = leftChildIndex;
      }

      if (rightChildIndex < _heap.length &&
          _heap[rightChildIndex] < _heap[smallestChildIndex]) {
        smallestChildIndex = rightChildIndex;
      }

      if (smallestChildIndex != index) {
        _swap(index, smallestChildIndex);
        index = smallestChildIndex;
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
    print("Min Heap: $_heap");
  }
}

void main() {
  MinHeap minHeap = MinHeap();

  minHeap.insert(3);
  minHeap.insert(1);
  minHeap.insert(4);
  minHeap.insert(2);

  print("Original Min Heap:");
  minHeap.display();

  print("\nExtract Min Value: ${minHeap.extractMin()}");

  print("Min Heap after Extraction:");
  minHeap.display();
}
