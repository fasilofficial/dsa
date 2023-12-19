// merge_sort.dart

void merge(List<int> arr, int left, int middle, int right) {
  int n1 = middle - left + 1;
  int n2 = right - middle;

  List<int> leftArray = List<int>.filled(n1, 0);
  List<int> rightArray = List<int>.filled(n2, 0);

  for (int i = 0; i < n1; i++) {
    leftArray[i] = arr[left + i];
  }

  for (int j = 0; j < n2; j++) {
    rightArray[j] = arr[middle + 1 + j];
  }

  int i = 0, j = 0, k = left;

  while (i < n1 && j < n2) {
    if (leftArray[i] <= rightArray[j]) {
      arr[k] = leftArray[i];
      i++;
    } else {
      arr[k] = rightArray[j];
      j++;
    }
    k++;
  }

  while (i < n1) {
    arr[k] = leftArray[i];
    i++;
    k++;
  }

  while (j < n2) {
    arr[k] = rightArray[j];
    j++;
    k++;
  }
}

void mergeSort(List<int> arr, int left, int right) {
  if (left < right) {
    int middle = (left + right) ~/ 2;

    mergeSort(arr, left, middle);
    mergeSort(arr, middle + 1, right);

    merge(arr, left, middle, right);
  }
}

void main() {
  List<int> array = [64, 34, 25, 12, 22, 11, 90];
  print("Original array: $array");

  mergeSort(array, 0, array.length - 1);

  print("Sorted array: $array");
}
