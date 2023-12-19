// graph.dart

class Graph {
  Map<String, List<String>> _adjacencyList;

  Graph() : _adjacencyList = {};

  void addVertex(String vertex) {
    if (!_adjacencyList.containsKey(vertex)) {
      _adjacencyList[vertex] = [];
    }
  }

  void addEdge(String vertex1, String vertex2) {
    _adjacencyList[vertex1]?.add(vertex2);
    _adjacencyList[vertex2]?.add(vertex1);
  }

  void display() {
    for (var vertex in _adjacencyList.keys) {
      String neighbors = _adjacencyList[vertex]!.join(', ');
      print('$vertex -> $neighbors');
    }
  }
}

void main() {
  Graph myGraph = Graph();

  myGraph.addVertex('A');
  myGraph.addVertex('B');
  myGraph.addVertex('C');
  myGraph.addVertex('D');

  myGraph.addEdge('A', 'B');
  myGraph.addEdge('B', 'C');
  myGraph.addEdge('C', 'D');
  myGraph.addEdge('D', 'A');

  print("Graph Structure:");
  myGraph.display();
}
