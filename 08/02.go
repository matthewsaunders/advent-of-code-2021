package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func StringToStringSlice(s string) []string {
	var str []string
	for _, runeValue := range s { 
		str = append(str, string(runeValue))
	}
	return str
}

func SliceIntersection(a, b []string) []string {
	hash := make(map[string]bool)
	for _, e := range a {
		hash[e] = false
	}
	for _, e := range b {
		if _, found := hash[e]; found {
			hash[e] = true
		}
	}

	elems := []string{}
	for k, v := range hash {
		if v {
			elems = append(elems, k)
		}
	}

	return elems
}

func SliceDifference(a, b []string) []string {
	hash := make(map[string]bool)
	for _, e := range a {
		hash[e] = true
	}
	for _, e := range b {
		if _, found := hash[e]; found {
			hash[e] = false
		} else {
			hash[e] = true
		}
	}

	elems := []string{}
	for k, v := range hash {
		if v {
			elems = append(elems, k)
		}
	}

	return elems
}

func SliceSubtract(a, b []string) []string {
	hash := make(map[string]bool)
	for _, e := range a {
		hash[e] = true
	}
	for _, e := range b {
		if _, found := hash[e]; found {
			hash[e] = false
		}
	}

	elems := []string{}
	for k, v := range hash {
		if v {
			elems = append(elems, k)
		}
	}

	return elems
}

func SliceContains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func SliceEqual(a, b []string) bool {
	if (len(a) != len(b)) { return false }
	
	for _, e := range b {
		if (!SliceContains(a, e)) { return false }
	}

	return true
}

func NewSegmentsMap() map[string]string {
	segmentsMap := make(map[string]string)
	segmentsMap["a"] = "0"
	segmentsMap["b"] = "0"
	segmentsMap["c"] = "0"
	segmentsMap["d"] = "0"
	segmentsMap["e"] = "0"
	segmentsMap["f"] = "0"
	segmentsMap["g"] = "0"
	return segmentsMap
}

func CloneNumbers(orig []*Number) []*Number {
	clones := []*Number{}

	for i := 0; i < len(orig); i++ {
		clones = append(clones, &Number{
			segments: orig[i].segments,
		})
	}

	return clones
}

type Number struct {
	segments []string
}

func decipherLine(line string) int {
	parts := strings.Split(line, " | ")
	left, right := parts[0], parts[1]
	inputs := strings.Fields(left)
	outputs := strings.Fields(right)

	// Initialize this whole shindig
	segmentsMap := NewSegmentsMap()
	numbers := [10]*Number{}
	fiveSegmentNumbers := []*Number{}
	sixSegmentNumbers := []*Number{}

	for j := 0; j < len(inputs); j++ {
		number := new(Number)
		segments := StringToStringSlice(inputs[j])
		number.segments = segments

		switch len(number.segments) {
		case 2:
			numbers[1] = number
		case 3:
			numbers[7] = number
		case 4:
			numbers[4] = number
		case 5:
			fiveSegmentNumbers = append(fiveSegmentNumbers, number)
		case 6:
			sixSegmentNumbers = append(sixSegmentNumbers, number)
		case 7:
			numbers[8] = number
		}
	}

	// Start working through the algorithm...

	// Step 1: Find "a" mapping
	segmentsMap["a"] = SliceDifference(numbers[7].segments, numbers[1].segments)[0]

	// Step 2: Find number 6
	findSix:
	for i := 0; i < len(sixSegmentNumbers); i++ {
		segments := sixSegmentNumbers[i].segments
		result := SliceSubtract(segments, numbers[1].segments)
		if (len(result) == 5) {
			numbers[6] = sixSegmentNumbers[i]
			break findSix
		}
	}

	// Step 3: Determine c and f
	if (SliceContains(numbers[6].segments, numbers[1].segments[0])) {
		segmentsMap["f"] = numbers[1].segments[0]
		segmentsMap["c"] = numbers[1].segments[1]
	} else {
		segmentsMap["f"] = numbers[1].segments[1]
		segmentsMap["c"] = numbers[1].segments[0]
	}

	// Step 4: Find number 3
	subtract := []string{segmentsMap["a"], segmentsMap["c"], segmentsMap["f"]}
	findThree:
	for i := 0; i < len(fiveSegmentNumbers); i++ {
		segments := fiveSegmentNumbers[i].segments
		result := SliceSubtract(segments, subtract)
		if (len(result) == 2) {
			numbers[3] = fiveSegmentNumbers[i]
			break findThree
		}
	}

	// Step 5: Determine d
	segmentsMap["d"] = SliceIntersection(SliceSubtract(numbers[3].segments, subtract), SliceSubtract(numbers[4].segments, subtract))[0]

	// Step 6: Find number 9
	subtract = append(subtract, segmentsMap["d"])
	findNine:
	for i := 0; i < len(sixSegmentNumbers); i++ {
		segments := sixSegmentNumbers[i].segments
		result := SliceSubtract(segments, subtract)
		if (len(result) == 2) {
			numbers[9] = sixSegmentNumbers[i]
			break findNine
		}
	}

	// Step 7: Determine e
	segmentsMap["e"] = SliceSubtract(numbers[8].segments, numbers[9].segments)[0]

	// Step 8: Determine b
	segmentsMap["b"] = SliceSubtract(numbers[4].segments, []string{segmentsMap["c"], segmentsMap["d"], segmentsMap["f"]})[0]

	// Step 9: Determine g
	segmentsMap["g"] = SliceSubtract(numbers[8].segments, []string{segmentsMap["a"], segmentsMap["b"], segmentsMap["c"], segmentsMap["d"], segmentsMap["e"], segmentsMap["f"]})[0]

	fmt.Printf("  Segment Map:\n")
	for k, v := range segmentsMap {
		fmt.Printf("    %v -> %v\n", k, v)
	}

	translatedDictionary := TranslatedNumberSegments(segmentsMap)
	decipheredStr := ""

	fmt.Println("==========================")
	fmt.Printf("Outputs:\n")
	for i := 0; i < len(outputs); i++ {
		segments := StringToStringSlice(outputs[i])
		fmt.Printf("--token: %v\n", outputs[i])
		fmt.Printf("--slice: %v\n", segments)

		findMatch:
		for j, translatedSegments := range translatedDictionary {
			fmt.Printf(" %v == %v\n", segments, translatedSegments)
			if (SliceEqual(segments, translatedSegments)) {
				fmt.Printf("equal: %v\n", j)
				fmt.Printf(" %v == %v\n", segments, translatedSegments)
				decipheredStr += strconv.Itoa(j)
				break findMatch
			}
		}

	}

	fmt.Printf("decipheredStr: %v\n", decipheredStr)
	decipheredVal, _ := strconv.Atoi(decipheredStr)
	fmt.Println("==========================")

	return decipheredVal
}

func NumberSegments() [][]string {
	return [][]string {
		{"a", "b", "c", "e", "f", "g"},
		{"c", "f"},
		{"a", "c", "d", "e", "g"},
		{"a", "c", "d", "f", "g"},
		{"b", "c", "d", "f"},
		{"a", "b", "d", "f", "g"},
		{"a", "b", "d", "e", "f", "g"},
		{"a", "c", "f"},
		{"a", "b", "c", "d", "e", "f", "g"},
		{"a", "b", "c", "d", "f", "g"},
	}
}

func TranslatedNumberSegments(translationMap map[string]string) [][]string {
	translatedSegments := [][]string{}
	numberSegments := NumberSegments()

	for i := 0; i < len(numberSegments); i++ {
		segments := []string{}
		for j := 0; j < len(numberSegments[i]); j++ {
			segments = append(segments, translationMap[numberSegments[i][j]])
		}
		translatedSegments = append(translatedSegments, segments)
	}

	return translatedSegments
}

func main() {
	f, err := os.Open("08/input.txt")
	if err != nil {
		fmt.Println(err)
		panic("Error reading input file")
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	lines := []string{}
	for s.Scan() {
		lines = append(lines, s.Text())
	}

	sum := 0

	for i := 0; i < len(lines); i++ {
		decoded := decipherLine(lines[i])
		sum += decoded
		fmt.Printf("Line %v: %v\n", i, lines[i])
		fmt.Printf("  decoded: %v\n", decoded)
	}

	fmt.Printf("Sum: %v\n", sum)
}