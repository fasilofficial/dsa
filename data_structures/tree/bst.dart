// binary_search_tree.dart

class Node {
  dynamic data;
  Node? left;
  Node? right;

  Node(this.data) : left = null, right = null;
}

class BinarySearchTree {
  Node? root;

  BinarySearchTree() : root = null;

  void insert(dynamic data) {
    root = _insertRecursive(root, data);
  }

  Node? _insertRecursive(Node? root, dynamic data) {
    if (root == null) {
      return Node(data);
    }

    if (data < root.data) {
      root.left = _insertRecursive(root.left, data);
    } else if (data > root.data) {
      root.right = _insertRecursive(root.right, data);
    }

    return root;
  }

  bool search(dynamic data) {
    return _searchRecursive(root, data);
  }

  bool _searchRecursive(Node? root, dynamic data) {
    if (root == null) {
      return false;
    }

    if (data == root.data) {
      return true;
    } else if (data < root.data) {
      return _searchRecursive(root.left, data);
    } else {
      return _searchRecursive(root.right, data);
    }
  }

  void inOrderTraversal() {
    _inOrderTraversalRecursive(root);
  }

  void _inOrderTraversalRecursive(Node? root) {
    if (root != null) {
      _inOrderTraversalRecursive(root.left);
      print(root.data);
      _inOrderTraversalRecursive(root.right);
    }
  }
}

void main() {
  BinarySearchTree bst = BinarySearchTree();

  bst.insert(5);
  bst.insert(3);
  bst.insert(7);
  bst.insert(2);
  bst.insert(4);
  bst.insert(6);

  print("In-Order Traversal:");
  bst.inOrderTraversal();

  print("\nSearch for 4: ${bst.search(4)}");
  print("Search for 8: ${bst.search(8)}");
}
