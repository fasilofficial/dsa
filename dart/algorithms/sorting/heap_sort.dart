// heap_sort.dart

void buildHeap(List<int> arr, int n, int i) {
  int largest = i;
  int left = 2 * i + 1;
  int right = 2 * i + 2;

  if (left < n && arr[left] > arr[largest]) {
    largest = left;
  }

  if (right < n && arr[right] > arr[largest]) {
    largest = right;
  }

  if (largest != i) {
    int temp = arr[i];
    arr[i] = arr[largest];
    arr[largest] = temp;

    buildHeap(arr, n, largest);
  }
}

void heapify(List<int> arr, int n) {
  for (int i = n ~/ 2 - 1; i >= 0; i--) {
    buildHeap(arr, n, i);
  }
}

void heapSort(List<int> arr) {
  int n = arr.length;

  heapify(arr, n);

  for (int i = n - 1; i >= 0; i--) {
    int temp = arr[0];
    arr[0] = arr[i];
    arr[i] = temp;

    buildHeap(arr, i, 0);
  }
}

void main() {
  List<int> array = [64, 34, 25, 12, 22, 11, 90];
  print("Original array: $array");

  heapSort(array);

  print("Sorted array: $array");
}
