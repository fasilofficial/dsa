// linear_search.dart

int linearSearch(List<int> array, int target) {
  for (int i = 0; i < array.length; i++) {
    if (array[i] == target) {
      return i; // Return the index if the target is found
    }
  }
  return -1; // Return -1 if the target is not found
}

void main() {
  List<int> array = [2, 4, 6, 8, 10, 12, 14];
  int target = 8;

  int result = linearSearch(array, target);

  if (result != -1) {
    print("Element $target found at index $result.");
  } else {
    print("Element $target not found in the array.");
  }
}
