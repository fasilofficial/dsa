// factorial.dart

int factorial(int n) {
  if (n == 0 || n == 1) {
    return 1;
  }
  return n * factorial(n - 1);
}

void main() {
  int n = 5;
  print("Factorial of $n is: ${factorial(n)}");
}
