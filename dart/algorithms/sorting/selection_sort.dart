// selection_sort.dart

void selectionSort(List<int> arr) {
  int n = arr.length;

  for (int i = 0; i < n - 1; i++) {
    int minIndex = i;
    for (int j = i + 1; j < n; j++) {
      if (arr[j] < arr[minIndex]) {
        minIndex = j;
      }
    }

    // swap arr[i] and arr[minIndex]
    int temp = arr[i];
    arr[i] = arr[minIndex];
    arr[minIndex] = temp;
  }
}

void main() {
  List<int> array = [64, 34, 25, 12, 22, 11, 90];
  print("Original array: $array");
  
  selectionSort(array);
  
  print("Sorted array: $array");
}
