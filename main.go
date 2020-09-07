package main

import (
	"fmt"
	"flag"
	"log"
	"strconv"
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/lionsoul2014/ip2region/binding/golang/ip2region"
)

var (
	region *ip2region.Ip2Region
	r *gin.Engine
	host string
	port int
	database string
)

func Init(){
	var err error
	region, err = ip2region.New(database)
	if err != nil {
		log.Fatal(err)
	}
	gin.DisableConsoleColor()
	r = gin.Default()

	route()
}

func Close(){
	defer region.Close()
}

func ipFunc(c *gin.Context) {
	ip := c.DefaultQuery("ip", c.ClientIP())
	var regionData map[string]interface{}
	ret, err := region.MemorySearch(ip)
	if err != nil {
		c.JSON(400, gin.H{
			"err": err,
		})
		return
	}
	json_ret,_ := json.Marshal(&ret)
	json.Unmarshal(json_ret,&regionData)
	regionData["ip"] = ip
	c.JSON(200, regionData)
}

func route(){
	r.GET("/ip",ipFunc )
	r.Run(fmt.Sprintf("%s:%s",host,strconv.Itoa(port)))
}

func main() {
	flag.StringVar(&database,"db","ip2region.db","run with database file")
	flag.StringVar(&host,"h","0.0.0.0","run in host ip")
	flag.IntVar(&port,"p",80,"run in port")
	flag.Parse()
	Init()
	Close()
}
