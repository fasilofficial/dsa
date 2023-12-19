// fibonacci.dart

int fibonacci(int n) {
  if (n <= 1) {
    return n;
  }

  int fib = 1;
  int prevFib = 1;

  for (int i = 2; i < n; i++) {
    int temp = fib;
    fib += prevFib;
    prevFib = temp;
  }

  return fib;
}

void main() {
  int n = 10;
  print("Fibonacci sequence up to $n terms:");
  for (int i = 0; i < n; i++) {
    print(fibonacci(i));
  }
}
