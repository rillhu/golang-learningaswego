package main

import (
	"encoding/json"
	"fmt"
)



type UpdatePart struct {
	Loc string
	Utc string
}

type BasicPart struct {
	City string
	Cnty string
	Id string
	Lat string
	Lon string
	Update UpdatePart
}

type BasicWeather struct {
	Basic BasicPart
	Now string
	Status string
}

type HeWeatherSlice struct {
	HeWeather5 []BasicWeather
}


//已知JSON格式的解析方法：定义与json数据对应的结构体，数组对应slice，字段名对应JSON里面的KEY,
//当你接收到一个很大的JSON数据结构而你却只想获取其中的部分数据的时候，你只需将你想要的数据对应的字段名大写,
//这是因为能够被赋值的字段必须是可导出字段(即首字母大写）。同时JSON解析的时候只会解析能找得到的字段，找不到的字段会被忽略
func main() {

	var w HeWeatherSlice
	str := `{"HeWeather5":[{"basic":{"city":"yanta","cnty":"China","id":"CN101110113","lat":"34.21339035","lon":"108.92658997","update":{"loc":"2018-04-11 08:47","utc":"2018-04-11 00:47"}},"now":{"cond":{"code":"101","txt":"Cloudy"},"fl":"15","hum":"46","pcpn":"0.0","pres":"1010","tmp":"17","vis":"4","wind":{"deg":"154","dir":"SE","sc":"1","spd":"3"}},"status":"ok"}]}`

	json.Unmarshal([]byte(str), &w)
	fmt.Println(w) //打印全部

	//打印部分字段
	fmt.Println(w.HeWeather5[0].Basic.City)
	fmt.Println(w.HeWeather5[0].Status)
}