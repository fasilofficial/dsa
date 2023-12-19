// hash_table.dart

class HashTable {
  late List<List<MapEntry<String, dynamic>>> _buckets;
  late int _size;

  HashTable(int size) {
    _size = size;
    _buckets = List.generate(size, (index) => []);
  }

  int _hashFunction(String key) {
    int hash = 0;
    for (int i = 0; i < key.length; i++) {
      hash = (hash + key.codeUnitAt(i)) % _size;
    }
    return hash;
  }

  void insert(String key, dynamic value) {
    int index = _hashFunction(key);
    _buckets[index].add(MapEntry(key, value));
  }

  dynamic get(String key) {
    int index = _hashFunction(key);
    for (var entry in _buckets[index]) {
      if (entry.key == key) {
        return entry.value;
      }
    }
    return null; // Key not found
  }

  void remove(String key) {
    int index = _hashFunction(key);
    _buckets[index].removeWhere((entry) => entry.key == key);
  }

  void display() {
    for (int i = 0; i < _size; i++) {
      print("Bucket $i:");
      for (var entry in _buckets[i]) {
        print("${entry.key}: ${entry.value}");
      }
    }
  }
}

void main() {
  HashTable myHashTable = HashTable(5);

  myHashTable.insert("name", "John");
  myHashTable.insert("age", 25);
  myHashTable.insert("city", "New York");

  print("Original Hash Table:");
  myHashTable.display();

  print("\nValue for 'name': ${myHashTable.get("name")}");

  myHashTable.remove("age");

  print("\nHash Table after removal:");
  myHashTable.display();
}
