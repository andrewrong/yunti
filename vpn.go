package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	ip := "ip.txt"

	file, err := os.Open(ip)
	defer file.Close()

	if err != nil {
		fmt.Println("open file is error,error:%v", err)
		return
	}

	br := bufio.NewReader(file)
	services := make(map[string]float64)

	for line, _, err := br.ReadLine(); err != io.EOF; line, _, err = br.ReadLine() {
		split := strings.Split(string(line), ":")
		fmt.Printf("name:%s,ip:%s \n", split[0], split[1])

		//执行命令
		cmd := exec.Command("ping", "-c", "4", split[1])
		buf, err := cmd.Output()
		if err != nil {
			fmt.Printf("执行ping出错,error:%v \n", err)
			continue
		}
		pattern := regexp.MustCompile(`time=(\d+.\d+)`)

		var result []string = pattern.FindAllString(string(buf), -1)

		count := 0
		sumTime := 0.0
		for _, time := range result {
			count++
			tmp := strings.Split(time, "=")
			floatValue, _ := strconv.ParseFloat(tmp[1], 64)
			sumTime += floatValue
		}

		services[split[0]] = (sumTime) / float64(count)
	}

	for name, time := range services {
		fmt.Printf("name:%s,time:%f\n", name, time)
	}
}
