package main

import (    
    "github.com/influxdata/influxdb1-client/v2"
    "time"
    "errors"
)

const (
    USERNAME string = "admin"
    PASSWORD string = "test12345"
    DATABASE string = "iot_db"
)

func createPoint(params map[string]interface{}) (map[string]interface{}, error) { 

    influxClient, err := client.NewHTTPClient(client.HTTPConfig{
              Addr: "http://localhost:8086",
              Username: USERNAME,    
              Password: PASSWORD,
              })

    if err != nil {
        return nil, err
    }

    bp, err := client.NewBatchPoints(client.BatchPointsConfig{    
    Database: DATABASE,
})

    if err != nil {
        return nil, err
    } 

    clientId := params["gpu_id"].(string) 
    location := params["rig_id"].(string)
    currentWatts := params["current_watts"]

    pt, err := client.NewPoint("total_watts", map[string]string{"gpu_id": gpuId, "rig_id": rigId},
                     map[string]interface{}{"current_watts": currentWatts},
                     time.Now())

    if err != nil {
        return nil, err
    } 

    bp.AddPoint(pt)

    err =  influxClient.Write(bp)         

    if err != nil {
        return nil, err           
    }

    resp := map[string]interface{}{"data" : "Added"}
    return resp, nil            
}

func getPoints() (map[string]interface{}, error) { 

    influxClient, err := client.NewHTTPClient(client.HTTPConfig{
              Addr: "http://localhost:8086",
              Username: USERNAME,    
              Password: PASSWORD,
              })

    if err != nil {
        return nil, err
    }

    queryString := "SELECT gpu_id, rig_id, current_watts FROM total_watts"
    q := client.NewQuery(queryString, DATABASE, "ns")
    response, err := influxClient.Query(q)

    if err != nil {
        return nil, err
    }
    err = response.Error()
if err != nil {
         return nil, errors.New("Empty data not allowed")
} else {

        res := response.Results
        if (len(res) == 0) {
           return nil, err
        }

        columns := response.Results[0].Series[0].Columns
        points  := response.Results[0].Series[0].Values              

        data := []map[string]interface{}{}                               

        for i := 0; i <= len(points) - 1 ; i++ {   

            record  := map[string]interface{}{}                  

            for j := 0; j <= len(columns) - 1; j++ {                   
                record[string(columns[j])] = points[i][j]                   
            }

            data = append(data, record)   

        }

        resp := map[string]interface{}{"data" : data}

        return resp, nil    
    }
}
