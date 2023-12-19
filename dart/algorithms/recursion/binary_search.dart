// binary_search_recursive.dart

int binarySearch(List<int> array, int target, int left, int right) {
  if (left <= right) {
    int mid = (left + right) ~/ 2;

    if (array[mid] == target) {
      return mid; // Return the index if the target is found
    } else if (array[mid] < target) {
      return binarySearch(array, target, mid + 1, right);
    } else {
      return binarySearch(array, target, left, mid - 1);
    }
  }

  return -1; // Return -1 if the target is not found
}

void main() {
  List<int> array = [2, 4, 6, 8, 10, 12, 14];
  int target = 8;

  int result = binarySearch(array, target, 0, array.length - 1);

  if (result != -1) {
    print("Element $target found at index $result.");
  } else {
    print("Element $target not found in the array.");
  }
}
