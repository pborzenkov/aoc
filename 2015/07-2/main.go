package main

import (
	"bufio"
	"log"
	"regexp"
	"strconv"

	"github.com/pborzenkov/aoc/pkg/input"
)

var (
	reAssignOp = regexp.MustCompile("^([[:alnum:]]+) -> ([[:alpha:]]+)$")
	reBinaryOp = regexp.MustCompile("^([[:alnum:]]+) (AND|OR) ([[:alnum:]]+) -> ([[:alpha:]]+)$")
	reShiftOp  = regexp.MustCompile("^([[:alpha:]]+) (LSHIFT|RSHIFT) ([[:digit:]]+) -> ([[:alpha:]]+)$")
	reNotOp    = regexp.MustCompile("^NOT ([[:alpha:]]+) -> ([[:alpha:]]+)$")

	nodes = make(map[string]func() int)
	cache = make(map[string]int)
)

func addNode(n string, fn func() int) {
	if _, ok := nodes[n]; ok {
		log.Fatalf("Node %q already exists", n)
	}
	nodes[n] = fn
}

func addNodeOverride(n string, fn func() int) {
	if _, ok := nodes[n]; !ok {
		log.Fatalf("Node %q doesn't exist", n)
	}
	nodes[n] = fn
}

func getNodeVal(n string) int {
	if val, ok := cache[n]; ok {
		return val
	}
	fn, ok := nodes[n]
	if !ok {
		log.Fatalf("Node %q doesn't exist", n)
	}
	val := fn()
	cache[n] = val
	return val
}

func val(n string) int {
	val, err := strconv.ParseInt(n, 10, 32)
	if err != nil {
		return getNodeVal(n)
	}
	return int(val)
}

func main() {
	r := input.NewFileOrStdin()
	s := bufio.NewScanner(r)
	for s.Scan() {
		if groups := reAssignOp.FindStringSubmatch(s.Text()); groups != nil {
			addNode(groups[2], func() int {
				return val(groups[1])
			})
			continue
		}
		if groups := reBinaryOp.FindStringSubmatch(s.Text()); groups != nil {
			addNode(groups[4], func() int {
				l, r := val(groups[1]), val(groups[3])
				switch groups[2] {
				case "AND":
					return l & r
				case "OR":
					return l | r
				default:
					log.Fatalf("Unknown binary operation %q", groups[2])
					return 0
				}
			})
			continue
		}
		if groups := reShiftOp.FindStringSubmatch(s.Text()); groups != nil {
			addNode(groups[4], func() int {
				shift, err := strconv.ParseUint(groups[3], 10, 32)
				if err != nil {
					log.Fatalf("Failed to parse %q: %v", groups[3], err)
				}
				l := getNodeVal(groups[1])
				switch groups[2] {
				case "LSHIFT":
					return l << uint(shift)
				case "RSHIFT":
					return l >> uint(shift)
				default:
					log.Fatalf("Unknown shift operation %q", groups[3])
					return 0
				}
			})
			continue
		}
		if groups := reNotOp.FindStringSubmatch(s.Text()); groups != nil {
			addNode(groups[2], func() int { return ^getNodeVal(groups[1]) })
			continue
		}
		log.Fatalf("unknown command: %q", s.Text())
	}
	if s.Err() != nil {
		log.Fatalf("Scanning completed with error (%v), result might be unreliable", s.Err())
	}
	a := getNodeVal("a")
	addNodeOverride("b", func() int { return a })
	for k := range cache {
		delete(cache, k)
	}
	log.Printf("Value of node 'a': %d", getNodeVal("a"))
}
