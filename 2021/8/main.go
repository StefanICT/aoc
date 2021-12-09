package main

import (
    . "aoc/util"
    "fmt"
    "strings"
    "strconv"
)

func main() {
	part1()
	part2()
}

func part1() {
    sum := 0
    for _, line := range ReadLines("2021/8/input") {
        components := strings.Fields(strings.Split(line, "|")[1])
        for _, component := range components {
            length := len(component)

            if length == 2 || length == 4 || length == 3 || length == 7 {
                sum++
            }
        }
    }

	fmt.Printf("Day 6.1: %d\n", sum)
}

func Used(data map[string][]string, needle string) bool {
    for _, values := range data {
        if Contains(values, needle) {
            return true
        }
    }

    return false
}

func Contains(values []string, needle string) bool {
    for _, value := range values {
        if value == needle {
            return true
        }
    }

    return false 
}

func Compare(s string, o string) bool {
    if len(s) != len(o) {
        return false
    }

    next:
    for _, c := range s {
        for _, c2 := range o {
            if c2 == c {
                continue next
            }
        }

        return false
    }
    lext:
    for _, c := range o {
        for _, c2 := range s {
            if c2 == c {
                continue lext
            }
        }

        return false
    }

    return true
}


func Remove(s [][]string, i int) [][]string { 
    return append(s[:i], s[i+1:]...)
}

func part2() {
    sum := 0
    for _, line := range ReadLines("2021/8/input") {
        sum += run(line)
    }
    fmt.Println(sum)
}

func run(line string) int {
    components := strings.Split(line, " | ")
    signals := strings.Fields(components[0])
    values := strings.Fields(components[1])

    blueprints := map[int][]string {
        0: []string{ "a", "b", "c",  "e", "f", "g" },
        1: []string{ "c", "f" },
        2: []string{ "a", "c", "d", "e", "g" },
        3: []string{ "a", "c", "d", "f", "g" },
        4: []string{ "b", "c", "d", "f" },
        5: []string{ "a", "b", "d", "f", "g" },
        6: []string{ "a", "b", "d", "e", "f", "g" },
        7: []string{ "a", "c", "f" },
        8: []string{ "a", "b", "c", "d", "e", "f", "g" },
        9: []string{ "a", "b", "c", "d", "f", "g" },
    }

    data := map[string][]string{}

    for _, i := range []int{2, 3, 4, 7 } {
        matches := []string{}
        for _, signal := range signals {
            if len(signal) == i {
                matches = append(matches, signal)
            }
        }

        var blueprint []string
        for _, bp := range blueprints {
            if len(bp) == i {
                blueprint = bp
                break
            }
        }

        tmp := map[string][]string{}

        for _, match := range matches {
            for _, character := range strings.Split(match, "") {
                for _, segment := range blueprint {
                    if !Used(data, character) && len(data[segment]) == 0 {
                        tmp[segment] = append(tmp[segment], character)
                    }
                }
            }
        }

        for key, values := range tmp {
            for _, value := range values {
                data[key] = append(data[key], value)
            }
        }
    }

    x := map[string]int{}
    for _, i := range []int{5, 6, 2, 3, 4, 7} {
        matches := []string{}
        for _, signal := range signals {
            if len(signal) == i {
                matches = append(matches, signal)
            }
        }

        bps := map[int][]string{}
        for digit, bp := range blueprints {
            if len(bp) == i {
                bps[digit] = bp
            }
        }

        matches: 
        for _, match := range matches {
            for digit, blueprint := range bps {
                values := [][]string{}
                for _, segment := range blueprint {
                    values = append(values, data[segment])
                }

                for _, character := range strings.Split(match, "") {
                    if len(values) == 0 {
                        continue matches
                    }

                    for i := len(values) - 1; i >= 0; i-- {
                        v := values[i]
                        if Contains(v, character) {
                            values = Remove(values, i)
                            break
                        }
                    }

                }

                if len(values) == 0 {
                    x[match] = digit
                    continue matches
                } 
            }
        }

    }

    output := ""
    for _, value := range values {
        for l, digit := range x {
            if Compare(value, l) {
                output = output + strconv.Itoa(digit)
            }

        }
    }

    return (AsInt(output))
}

