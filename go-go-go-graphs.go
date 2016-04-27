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
*/


func (graph *Graph) Dijkstra(source string, target string)(map[string]int,map[string]*Node){
    
      //Unvisted
      Q := make(map[string]*Node)
      dist := make(map[string]int)
      prev := make(map[string]*Node)
      
      for key , val := range graph.Nodes {             // Initialization
          dist[key] = 9999999999                  // Unknown distance from source to v
          prev[key] = nil                 // Previous node in optimal path from source
          Q[key] = val                     // All nodes initially in Q (unvisited nodes)
      }
      
      dist[source] = 0                        // Distance from source to source
      
      for(len(Q) != 0){
          u ,keyy := min(Q,dist)    // Source node will be selected first
          
      
          //for each neighbor v of u: 
          for key, value := range Q[keyy].Edges{          // where v is still in Q.
          
              alt := dist[keyy] + value
              if alt < dist[key]{               // A shorter path to v has been found
                  dist[key] = alt 
                  prev[key] = u 
              }
          }
          delete(Q,keyy) //remove u from Q 
      }
      return dist, prev
    
    
    
}

//u := *Node in Q with min dist[u] 
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
   
    
    weight, prev := graph1.Dijkstra("A","E")
   
     //Turn our Router to json , this was for testing

    res1B, _ := json.Marshal(graph1)
    fmt.Println(string(res1B))
    fmt.Println("\n\n\n")
    
    res1B, _ = json.Marshal(weight)
    fmt.Println(string(res1B))
    
    fmt.Println("\n\n\n")
    
    res1B, _ = json.Marshal(prev)
    fmt.Println(string(res1B))

}

