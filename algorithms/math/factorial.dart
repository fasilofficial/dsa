// factorial.dart

int factorial(int n) {
  int result = 1;

  for (int i = 2; i <= n; i++) {
    result *= i;
  }

  return result;
}

void main() {
  int n = 5;
  print("Factorial of $n is: ${factorial(n)}");
}
