package main

import(
    "fmt"
    "encoding/json"
)

/*
 From: https://en.wikipedia.org/wiki/Dijkstra%27s_algorithm
 
 1  function Dijkstra(Graph, source):
 2
 3      create vertex set Q
 4
 5      for each vertex v in Graph:             // Initialization
 6          dist[v] ← INFINITY                  // Unknown distance from source to v
 7          prev[v] ← UNDEFINED                 // Previous node in optimal path from source
 8          add v to Q                          // All nodes initially in Q (unvisited nodes)
 9
10      dist[source] ← 0                        // Distance from source to source
11      
12      while Q is not empty:
13          u ← vertex in Q with min dist[u]    // Source node will be selected first
14          remove u from Q 
15          
16          for each neighbor v of u:           // where v is still in Q.
17              alt ← dist[u] + length(u, v)
18              if alt < dist[v]:               // A shorter path to v has been found
19                  dist[v] ← alt 
20                  prev[v] ← u 
21
22      return dist[], prev[]

Now we can read the shortest path from source to target by reverse iteration:

1  S ← empty sequence
2  u ← target
3  while prev[u] is defined:                  // Construct the shortest path with a stack S
4      insert u at the beginning of S         // Push the vertex onto the stack
5      u ← prev[u]                            // Traverse from target to source
6  insert u at the beginning of S             // Push the source onto the stack


*/

//func (graph *Graph) Dijkstra(source string, target string)(dist map[string]int,prev map[string]*Node){
func (graph *Graph) Dijkstra(source string, target string)(dist map[string]int,path []string){
    
      //Unvisted
      Q := make(map[string]*Node)
      dist = make(map[string]int)
      prev := make(map[string]*Node)
      
      for key , val := range graph.Nodes {             // Initialization
          dist[key] = 9999999999                  // Unknown distance from source to v
          prev[key] = nil                 // Previous node in optimal path from source
          Q[key] = val                     // All nodes initially in Q (unvisited nodes)
      }
      
      dist[source] = 0                        // Distance from source to source
      
      for(len(Q) != 0){
          u ,keyy := min(Q,dist)    // Source node will be selected first    
          delete(Q,keyy) //remove u from Q 
          //for each neighbor v of u: 
          for key, value := range u.Edges{          // where v is still in Q.
          
              alt := dist[keyy] + value
              if alt < dist[key]{               // A shorter path to v has been found
                  dist[key] = alt 
                  prev[key] = u 
              }
          }
          
      }
      
      
      /*
1  S ← empty sequence
2  u ← target
3  while prev[u] is defined:                  // Construct the shortest path with a stack S
4      insert u at the beginning of S         // Push the vertex onto the stack
5      u ← prev[u]                            // Traverse from target to source
6  insert u at the beginning of S             // Push the source onto the stack


*/
      
        var S []string
        u := target
        // Construct the shortest path with a stack S
        for(prev[u] != nil){
            // Push the vertex onto the stack
            S = append([]string{u},S...)
            u = graph.getKey(prev[u])                            // Traverse from target to source
        }
        
      return dist, S
    
    
    
}

//u := *Node in Q with min dist[u] 

func (graph *Graph) getKey(node *Node) (key string) {
    for key ,val := range graph.Nodes{
        if(val == node){
            return key
        }
    }
    
    return "ERROR"
}

func min(Q map[string]*Node,dist map[string]int) (*Node,string) {
    
    var min int
    var minNode *Node
    firstLoop := true
    var keyy string
    
    for key ,value := range Q{
        if(firstLoop){
            min = dist[key]
            minNode = value
            firstLoop = false
            keyy = key
        }else{
            if(dist[key]< min){
                min = dist[key]
                minNode = value    
                keyy = key   
            }
        }
    }
    
    
    return minNode ,keyy
}


type Node struct{
    Edges map[string]int
}

type Graph struct{
    Name string
    TotalNodes int
    Nodes map[string]*Node
}


func (graph *Graph) addUndirectedWeightedVertice(idOfNode1 string ,idOfNode2 string ,weight int){
    
    graph.Nodes[idOfNode1].Edges[idOfNode2] = weight
    graph.Nodes[idOfNode2].Edges[idOfNode1] = weight
    
}

func createGraph(graphName string) *Graph {
    graphThis := Graph{ 
        Name : graphName,
        TotalNodes : 0,
    }
    
    graphThis.Nodes = make(map[string]*Node)
    
    return &graphThis
}

func (graph *Graph) addNode(idOfNode string) {
    
    //TODO: Add Real error
    if _ , isKey := graph.Nodes[idOfNode]; isKey{
        fmt.Println("Key already exists. Did not add node")
    }else{
        
    
        graph.TotalNodes = graph.TotalNodes + 1
        newNode := Node{} 

        graph.Nodes[idOfNode] = &newNode

        newNode.Edges = make(map[string]int)
            
        
    }
    
}

func main()  {
    
    graph1 := createGraph("Fun Graph")
    
    graph1.addNode("A")
    graph1.addNode("B")
    graph1.addNode("C")
    graph1.addNode("D")
    graph1.addNode("E")
    
    graph1.addUndirectedWeightedVertice("A","B",2)
    graph1.addUndirectedWeightedVertice("C","A",3)
    graph1.addUndirectedWeightedVertice("C","B",2)
    graph1.addUndirectedWeightedVertice("B","E",18)
    graph1.addUndirectedWeightedVertice("B","D",1)
    graph1.addUndirectedWeightedVertice("C","D",2)
    graph1.addUndirectedWeightedVertice("D","E",9)
   
    
    weight, path := graph1.Dijkstra("A","E")
   
     //Turn our Router to json , this was for testing

    res1B, _ := json.Marshal(graph1)
    fmt.Println(string(res1B))
    fmt.Println("\n\n\n")
    
    res1B, _ = json.Marshal(weight)
    fmt.Println(string(res1B))
    
    fmt.Println("\n\n\n")
    
    res1B, _ = json.Marshal(path)
    fmt.Println(string(res1B))

}

