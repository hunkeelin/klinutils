package klinutils

import (
	"fmt"
	"testing"
)

func TestRandint(t *testing.T) {
	f := RandInt(0, 34)
	d := RandInt(0, 34)
	e := RandInt(0, 34)
	fmt.Println(f, d, e)
}
func TestPrint(t *testing.T) {
	fmt.Println(Stringtoport("superca"))
}
func TestWget(t *testing.T) {
	fmt.Println("testing wget")
	w := WgetInfo{
		Dest:  "util3.klin-pro.com",
		Dport: "46861",
		Route: "cacerts/rootca.crt",
	}
	b, err := Wget(w)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}
func TestAlgo(t *testing.T) {
	fmt.Println(Stringtoport("ssh"))
}
func TestGenv2(t *testing.T) {
	f, err := Genuuidv2("fuck", 3, 43)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(f))
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
