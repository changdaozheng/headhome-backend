package logic

import (
    "os"
    "fmt"
    "errors"
    "strings"
    "net/http"
    "io/ioutil"
    "encoding/json"

    "github.com/joho/godotenv"
)

type MultiTypes struct {
    Text  string `json:"text"`
    Value int64  `json:"value"`
}

type Location struct {
    Lat float64 `json:"lat"`
    Lng float64 `json:"lng"`
}

type Polyline struct {
    Points string `json:"points"`
}

type Node struct {
    Distance         MultiTypes `json:"distance"`
    Duration         MultiTypes `json:"duration"`
    StartLocation    Location `json:"start_location"`
    EndLocation      Location `json:"end_location"`
    HTMLInstructions string   `json:"html_instructions"`
    Maneuver         string   `json:"maneuver"`
    Polyline         Polyline `json:"polyline"`
    TravelMode       string   `json:"travel_mode"`
}

type NewNode struct {
    Distance         int64 `json:"Distance"`
    Duration         int64 `json:"Duration"`
    StartLocation    Location `json:"StartLocation"`
    EndLocation      Location `json:"EndLocation"`
    HTMLInstructions string   `json:"HTMLInstructions"`
    Maneuver         string   `json:"Maneuver"`
    Polyline         string `json:"Polyline"`
    TravelMode       string   `json:"TravelMode"`
}

type DirectionsResult struct {
    OverallPolyline    string    `json:"Polyline"`
    Route              []NewNode   `json:"Route"`
}

func mapToNode(data interface{}) (Node, error){
    //assert map[string]interface on input 
    m, ok := data.(map[string]interface{})
    
    if !ok {
        return Node{}, errors.New("nodes retrieved is of incorrect form")
    }

    //map to []byte
    jsonBytes, err := json.Marshal(m)
    if err != nil {
        return Node{}, err
    }

    //Bind []byte to Node 
    var node Node 
    if err := json.Unmarshal(jsonBytes, &node); err != nil {
        return Node{}, err
    }

    return node, nil
}

func RetrieveDirections(start string, end string) (DirectionsResult, error){

    //Step 1: Directions API call
    //Load API key from .env
    err := godotenv.Load()
	if err != nil {
		return DirectionsResult{}, errors.New("error loading .env file")
	}

    //API call inputs
    start = strings.Replace(start, " ", "+", -1)
    end = strings.Replace(end, " ", "+", -1)
    apiKey := os.Getenv("MAPS_API_KEY")
    url := fmt.Sprintf("https://maps.googleapis.com/maps/api/directions/json?origin=%s&destination=%s&mode=walking&key=%s", start, end, apiKey)
    method := "GET"
    
    client := &http.Client {
    }

    req, err := http.NewRequest(method, url, nil)
    if err != nil {
      return DirectionsResult{}, err 
    }
    
    res, err := client.Do(req)
    if err != nil {
        return DirectionsResult{}, err 
    }
    defer res.Body.Close()

    data, err := ioutil.ReadAll(res.Body)
    if err != nil {
        return DirectionsResult{}, err 
    }

    //Convert response body to JSON
    var jsonData map[string]interface{}
    if err = json.Unmarshal(data, &jsonData); err != nil {
        return DirectionsResult{}, err
    }
    
    

    //Step 2: Formatting relevant output

    /*
    Extracting relevant nested json data
    a. steps = res['routes'][0]['legs'][0]['steps']
    b. polyline = res['routes'][0]['overview_polyline']
    */
    routes := jsonData["routes"].([]interface{})
    firstRoute :=routes[0].(map[string]interface{})

    //Polyline data formatting
    polylineMap := firstRoute["overview_polyline"].(map[string]interface{})
    polyline := Polyline{
        Points: polylineMap["points"].(string),
    }

    //Route node data formatting
    nodes := firstRoute["legs"].([]interface{})[0].(map[string]interface{})["steps"].([]interface{})
    
    var nodeSlice []Node
    for _, node := range nodes {
        res, err := mapToNode(node)
        if err!= nil {
            return DirectionsResult{}, err
        }

        nodeSlice = append(nodeSlice, res)
    }

    var newNodeSlice []NewNode
    for _, node := range nodeSlice {
        newNode := NewNode {
            Distance: node.Distance.Value,
            Duration: node.Duration.Value,
            StartLocation: node.StartLocation,
            EndLocation: node.EndLocation,
            HTMLInstructions: node.HTMLInstructions,
            Maneuver: node.Maneuver,
            Polyline: node.Polyline.Points,
            TravelMode: node.TravelMode,
        }

        newNodeSlice = append(newNodeSlice, newNode)
    }

    //Format results 
    result := DirectionsResult {
        OverallPolyline: polyline.Points,
        Route: newNodeSlice,
    }

    return result, nil
}