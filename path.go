package main

import "strings"

type Path string

func (p Path) ToString() string {
	return string(p)
}

func (p Path) Extension() string {
	splittedPath := strings.Split(p.ToString(), ".")
	return splittedPath[len(splittedPath)-1]
}