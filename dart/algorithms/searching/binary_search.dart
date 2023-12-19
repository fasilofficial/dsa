// binary_search.dart

int binarySearch(List<int> array, int target) {
  int left = 0;
  int right = array.length - 1;

  while (left <= right) {
    int mid = (left + right) ~/ 2;

    if (array[mid] == target) {
      return mid; // Return the index if the target is found
    } else if (array[mid] < target) {
      left = mid + 1;
    } else {
      right = mid - 1;
    }
  }

  return -1; // Return -1 if the target is not found
}

void main() {
  List<int> array = [2, 4, 6, 8, 10, 12, 14];
  int target = 8;

  int result = binarySearch(array, target);

  if (result != -1) {
    print("Element $target found at index $result.");
  } else {
    print("Element $target not found in the array.");
  }
}
