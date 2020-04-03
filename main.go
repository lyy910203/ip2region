package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lionsoul2014/ip2region/binding/golang/ip2region"
)

var (
	region *ip2region.Ip2Region
	r *gin.Engine
	host string
	port int

)


func Init(){
	var err error
	region, err = ip2region.New("ip2region.db")
	if err != nil {
		fmt.Println(err)
		return
	}
	gin.DisableConsoleColor()
	r = gin.Default()

	route() // 加载路由
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
	fmt.Println()
	c.JSON(200, regionData)
}

func route(){
	r.GET("/ip",ipFunc )
	r.Run(fmt.Sprintf("%s:%s",host,strconv.Itoa(port)))
}
func main() {
	flag.StringVar(&host,"h","0.0.0.0","run in host ip,default 0.0.0.0")
	flag.IntVar(&port,"p",80,"run in port,default 80")
	flag.Parse()
	Init()
	Close()

}
