package main

import (
	"bufio"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/pborzenkov/aoc/pkg/input"
)

const (
	regA = iota
	regB
)

type state struct {
	ip   int
	regs [2]int
}

type Execer interface {
	Exec(*state)
}

type ExecerFunc func(*state)

func (f ExecerFunc) Exec(s *state) {
	f(s)
}

var (
	hlf = func(reg int) ExecerFunc {
		return ExecerFunc(func(s *state) {
			s.regs[reg] /= 2
			s.ip++
		})
	}
	tpl = func(reg int) ExecerFunc {
		return ExecerFunc(func(s *state) {
			s.regs[reg] *= 3
			s.ip++
		})
	}
	inc = func(reg int) ExecerFunc {
		return ExecerFunc(func(s *state) {
			s.regs[reg]++
			s.ip++
		})
	}
	jmp = func(off int) ExecerFunc {
		return ExecerFunc(func(s *state) {
			s.ip += off
		})
	}
	jie = func(reg int, off int) ExecerFunc {
		return ExecerFunc(func(s *state) {
			if s.regs[reg]%2 == 0 {
				s.ip += off
			} else {
				s.ip++
			}
		})
	}
	jio = func(reg int, off int) ExecerFunc {
		return ExecerFunc(func(s *state) {
			if s.regs[reg] == 1 {
				s.ip += off
			} else {
				s.ip++
			}
		})
	}
)

func register(r string) int {
	r = strings.Trim(r, " ")
	switch r {
	case "a":
		return regA
	case "b":
		return regB
	default:
		panic(fmt.Sprintf("unexpected register %q", r))
	}
}

func offset(o string) int {
	o = strings.Trim(o, " ")
	off, err := strconv.Atoi(o)
	if err != nil {
		panic(fmt.Sprintf("failed to parse jump offset %q: %v", o, err))
	}
	return off
}

func main() {
	r := input.NewFileOrStdin()
	s := bufio.NewScanner(r)

	var code []Execer
	var state state

	for s.Scan() {
		parts := strings.SplitN(s.Text(), " ", 2)
		switch strings.Trim(parts[0], " ") {
		case "hlf":
			code = append(code, hlf(register(parts[1])))
		case "tpl":
			code = append(code, tpl(register(parts[1])))
		case "inc":
			code = append(code, inc(register(parts[1])))
		case "jmp":
			code = append(code, jmp(offset(parts[1])))
		case "jie":
			operands := strings.Split(parts[1], ",")
			code = append(code, jie(register(operands[0]), offset(operands[1])))
		case "jio":
			operands := strings.Split(parts[1], ",")
			code = append(code, jio(register(operands[0]), offset(operands[1])))
		}
	}

	state.regs[regA] = 1
	for state.ip < len(code) {
		code[state.ip].Exec(&state)
	}

	log.Printf("Register b value is %d", state.regs[regB])
}
