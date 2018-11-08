package main

import (
"fmt"
"os"
"strconv"
"bufio"
"strings"
"sort"
)

func check(e error){
	if e!= nil{
		panic(e)}
}
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	/*********************************READ FILE************************************/
	inputfile := os.Args[1]			//get input file name
	f, err := os.Open(inputfile)
	check(err)
	defer 	f.Close()
	
	s := make([]string, 0)			//slice to hold all the strings read from the file
	line := make([]string, 0)
	cyl := make([]int, 0)
	tempcyl := make([]int, 0)
	flag := 0
	tcount := 0
	
	scanner:= bufio.NewScanner(f)
	for scanner.Scan(){	
		s = append(s, scanner.Text())
		//fmt.Println(scanner.Text())	//Print out the text being "read" in
	}
	/*******************************EXTRACT DATA***********************************/	
	line = strings.Split(s[0], " ") //Determine what type of scheduling is being used
	if(line[1] == "fcfs"){
		fmt.Println("Seek algorithm: FCFS")
		line = strings.Split(s[1], " ")
		lcyl, _ := strconv.Atoi(line[1])
		fmt.Printf("\tLower cylinder: %5d\n", lcyl)
		line = strings.Split(s[2], " ")
		ucyl, _ := strconv.Atoi(line[1])
		fmt.Printf("\tUpper cylinder: %5d\n", ucyl)
		line = strings.Split(s[3], " ")
		icyl, _  := strconv.Atoi(line[1])
		fmt.Printf("\tInit cylinder: %5d\n", icyl)
		fmt.Println("\tCylinder requests:")
		
		for i := 4; i < (len(s) - 1); i++{
			line = strings.Split(s[i], " ")
			temp, _ := strconv.Atoi(line[1])
			cyl = append(cyl, temp)
			fmt.Printf("\t\tCylinder %5d\n", temp)
		}
		for i := 0; i < len(cyl); i++{
			fmt.Printf("Servicing %5d\n", cyl[i])
		}
		
		temp := Abs(cyl[0] - icyl)  //first - init
		
		for i := 0; i < (len(cyl) - 1); i++{
			temp = temp + Abs(cyl[i+1] - cyl[i]) 
		}
		tcount := temp
		fmt.Printf("FCFS traversal count = %d\n", tcount)
		
	}else if(line[1] == "c-look"){
		fmt.Println("Seek algorithm: C-LOOK")
		line = strings.Split(s[1], " ")
		lcyl, _ := strconv.Atoi(line[1])
		fmt.Printf("\tLower cylinder: %5d\n", lcyl)
		line = strings.Split(s[2], " ")
		ucyl, _ := strconv.Atoi(line[1])
		fmt.Printf("\tUpper cylinder: %5d\n", ucyl)
		line = strings.Split(s[3], " ")
		icyl, _  := strconv.Atoi(line[1])
		fmt.Printf("\tInit cylinder: %5d\n", icyl)
		fmt.Println("\tCylinder requests:")
		
		for i := 4; i < (len(s) - 1); i++{
			line = strings.Split(s[i], " ")
			temp, _ := strconv.Atoi(line[1])
			cyl = append(cyl, temp)
			fmt.Printf("\t\tCylinder %5d\n", temp)
		}	
		sort.Ints(cyl)
		for i := 0; i < len(cyl); i++{
			if(cyl[i] > icyl){
			tempcyl = append(tempcyl, cyl[i])
				fmt.Printf("Servicing %d\n", cyl[i])
			}
		}
		for i := 0; i < len(cyl); i++{
			if(cyl[i] < icyl){
				tempcyl = append(tempcyl, cyl[i])
				fmt.Printf("Servicing %d\n", cyl[i])
			}
		}
		temp := Abs(tempcyl[0] - icyl)
		
		for i := 0; i < (len(tempcyl) - 1); i++{
			temp = temp + Abs(tempcyl[i+1] - tempcyl[i]) 
		}
		tcount := temp
		fmt.Printf("C-LOCK traversal count = %d\n", tcount)
		
	}else if(line[1] == "sstf"){
		fmt.Println("Seek algorithm: SSTF")
		line = strings.Split(s[1], " ")
		lcyl, _ := strconv.Atoi(line[1])
		fmt.Printf("\tLower cylinder: %5d\n", lcyl)
		line = strings.Split(s[2], " ")
		ucyl, _ := strconv.Atoi(line[1])
		fmt.Printf("\tUpper cylinder: %5d\n", ucyl)
		line = strings.Split(s[3], " ")
		icyl, _  := strconv.Atoi(line[1])
		fmt.Printf("\tInit cylinder: %5d\n", icyl)
		fmt.Println("\tCylinder requests:")
		
		for i := 4; i < (len(s) - 1); i++{
			line = strings.Split(s[i], " ")
			temp, _ := strconv.Atoi(line[1])
			cyl = append(cyl, temp)
			fmt.Printf("\t\tCylinder %5d\n", temp)
		}	
		sort.Ints(cyl)
		
		length := len(cyl)
		current := icyl
		for i := 0; i < length; i++{
			dist := Abs(current - cyl[0])
			index := 0
			
			for j := 1; j < len(cyl); j++{
				difference := Abs(current - cyl[j])
				if(difference < dist){
					dist = difference
					index = j
				}
			}
			current = cyl[index]
			tempcyl = append(tempcyl, cyl[index])
			fmt.Printf("Servicing %d\n", cyl[index])
			cyl = append(cyl[:index], cyl[index + 1:]...)
		}
		
		temp := Abs(tempcyl[0] - icyl)
		
		for i := 0; i < (len(tempcyl) - 1); i++{
			temp = temp + Abs(tempcyl[i+1] - tempcyl[i]) 
		}
		tcount := temp
		fmt.Printf("SSTF traversal count = %d\n", tcount) 

	}else if(line[1] == "scan"){
		fmt.Println("Seek algorithm: SCAN")
		line = strings.Split(s[1], " ")
		lcyl, _ := strconv.Atoi(line[1])
		fmt.Printf("\tLower cylinder: %5d\n", lcyl)
		line = strings.Split(s[2], " ")
		ucyl, _ := strconv.Atoi(line[1])
		fmt.Printf("\tUpper cylinder: %5d\n", ucyl)
		line = strings.Split(s[3], " ")
		icyl, _  := strconv.Atoi(line[1])
		fmt.Printf("\tInit cylinder: %5d\n", icyl)
		fmt.Println("\tCylinder requests:")
		
		for i := 4; i < (len(s) - 1); i++{
			line = strings.Split(s[i], " ")
			temp, _ := strconv.Atoi(line[1])
			cyl = append(cyl, temp)
			fmt.Printf("\t\tCylinder %5d\n", temp)
		}	
		sort.Ints(cyl)

		for i := 0; i < len(cyl); i++{
			if(cyl[i] > icyl){
				fmt.Printf("Servicing %d\n", cyl[i])
			}
		}
		for i := (len(cyl) - 1); i >= 0; i--{
			if(cyl[i] < icyl){
				flag = 1
				fmt.Printf("Servicing %d\n", cyl[i])
			}
		}
		if(flag == 1){
			tcount = 2*(ucyl - icyl) + Abs(icyl - cyl[0])		
		}else{
			tcount = cyl[len(cyl)-1] - icyl
		}

		fmt.Printf("SCAN traversal count = %d\n", tcount)
		
	}else if(line[1] == "c-scan"){
		fmt.Println("Seek algorithm: C-SCAN")
		line = strings.Split(s[1], " ")
		lcyl, _ := strconv.Atoi(line[1])
		fmt.Printf("\tLower cylinder: %5d\n", lcyl)
		line = strings.Split(s[2], " ")
		ucyl, _ := strconv.Atoi(line[1])
		fmt.Printf("\tUpper cylinder: %5d\n", ucyl)
		line = strings.Split(s[3], " ")
		icyl, _  := strconv.Atoi(line[1])
		fmt.Printf("\tInit cylinder: %5d\n", icyl)
		fmt.Println("\tCylinder requests:")	
		
		for i := 4; i < (len(s) - 1); i++{
			line = strings.Split(s[i], " ")
			temp, _ := strconv.Atoi(line[1])
			cyl = append(cyl, temp)
			fmt.Printf("\t\tCylinder %5d\n", temp)
		}	
		sort.Ints(cyl)

		for i := 0; i < len(cyl); i++{
			if(cyl[i] > icyl){
			tempcyl = append(tempcyl, cyl[i])
				fmt.Printf("Servicing %d\n", cyl[i])
			}
		}
		for i := 0; i < len(cyl); i++{
			if(cyl[i] < icyl){
				flag = 1
				tempcyl = append(tempcyl, cyl[i])
				fmt.Printf("Servicing %d\n", cyl[i])
			}
		}
		
		if(flag == 1){
			tcount = 2*(ucyl - icyl) + Abs(icyl + tempcyl[len(cyl) - 1])		
		}else{
			tcount = cyl[len(cyl)-1] - icyl
		}
		fmt.Printf("C-SCAN traversal count = %d\n", tcount)

	}else if(line[1] == "look"){
		fmt.Println("Seek algorithm: LOOK")
		line = strings.Split(s[1], " ")
		lcyl, _ := strconv.Atoi(line[1])
		fmt.Printf("\tLower cylinder: %5d\n", lcyl)
		line = strings.Split(s[2], " ")
		ucyl, _ := strconv.Atoi(line[1])
		fmt.Printf("\tUpper cylinder: %5d\n", ucyl)
		line = strings.Split(s[3], " ")
		icyl, _  := strconv.Atoi(line[1])
		fmt.Printf("\tInit cylinder: %5d\n", icyl)
		fmt.Println("\tCylinder requests:")
		
		for i := 4; i < (len(s) - 1); i++{
			line = strings.Split(s[i], " ")
			temp, _ := strconv.Atoi(line[1])
			cyl = append(cyl, temp)
			fmt.Printf("\t\tCylinder %5d\n", temp)
		}	
		
		sort.Ints(cyl)
		
		for i := 0; i < len(cyl); i++{
			if(cyl[i] > icyl){
			tempcyl = append(tempcyl, cyl[i])
				fmt.Printf("Servicing %d\n", cyl[i])
			}
		}
		for i := (len(cyl) - 1); i >= 0; i--{
			if(cyl[i] < icyl){
				tempcyl = append(tempcyl, cyl[i])
				fmt.Printf("Servicing %d\n", cyl[i])
			}
		}
		temp := Abs(tempcyl[0] - icyl)
		
		for i := 0; i < (len(tempcyl) - 1); i++{
			temp = temp + Abs(tempcyl[i+1] - tempcyl[i]) 
		}
		tcount := temp

		fmt.Printf("LOCK traversal count = %d\n", tcount)
	}
}