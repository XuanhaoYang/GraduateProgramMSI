package model

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

type Result struct {
	Code    int     `json:"code"`
	Weight  string  `json:"weight"`
	Value   string  `json:"value"`
	Message string  `json:"msg"`
	Score   float32 `json:"score"`
	Pre     float32 `json:"pre"`
}

func BulkDetection(sampName string) Result {
	sampPath := "./data/" + sampName
	resPath := "./result/" + sampName
	modelPath := "./python/model/bulk.m"
	pyPath := "./python/detection.py"
	fmt.Println("python", pyPath, "--model", modelPath, "--input", sampPath, "--output", resPath)
	out, _ := exec.Command("python", pyPath, "--model", modelPath, "--input", sampPath, "--output", resPath).Output()

	//反序列化
	var res Result
	err = json.Unmarshal(out, &res)
	if err != nil {
		fmt.Println("json unmarshal error:", err)
	}
	fmt.Println(res.Weight)
	fmt.Println(res.Value)
	return res
}
