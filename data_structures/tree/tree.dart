// tree.dart

class TreeNode {
  dynamic data;
  List<TreeNode> children;

  TreeNode(this.data) : children = [];

  void addChild(TreeNode child) {
    children.add(child);
  }

  void display(int depth) {
    String indentation = '  ' * depth;
    print('$indentation$data');
    for (var child in children) {
      child.display(depth + 1);
    }
  }
}

void main() {
  TreeNode root = TreeNode('Root');

  TreeNode child1 = TreeNode('Child 1');
  TreeNode child2 = TreeNode('Child 2');
  TreeNode child3 = TreeNode('Child 3');

  root.addChild(child1);
  root.addChild(child2);
  root.addChild(child3);

  TreeNode grandchild1 = TreeNode('Grandchild 1');
  child2.addChild(grandchild1);

  print("Tree Structure:");
  root.display(0);
}
