package klinutils

import (
	"fmt"
	"testing"
)

func TestWget(t *testing.T) {
	fmt.Println("testing wget")
	w := WgetInfo{
		Dest:  "util3.klin-pro.com",
		Dport: "2018",
		Route: "cacerts/rootca.crt",
	}
	b, _ := Wget(w)
	fmt.Println(string(b))
}
func TestAlgo(t *testing.T) {
	fmt.Println(Stringtoport("ssh"))
}
func TestGen(t *testing.T) {
	f, err := Genuuid()
	if err != nil {
		panic(err)
	}
	d, err := Genmac()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(f))
	fmt.Println(string(d))
}
