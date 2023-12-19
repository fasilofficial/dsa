// insertion_sort.dart

void insertionSort(List<int> arr) {
  int n = arr.length;

  for (int i = 1; i < n; ++i) {
    int key = arr[i];
    int j = i - 1;

    while (j >= 0 && arr[j] > key) {
      arr[j + 1] = arr[j];
      j = j - 1;
    }
    arr[j + 1] = key;
  }
}

void main() {
  List<int> array = [64, 34, 25, 12, 22, 11, 90];
  print("Original array: $array");
  
  insertionSort(array);
  
  print("Sorted array: $array");
}
