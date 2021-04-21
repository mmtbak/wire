package main

import (
	"fmt"
	"time"

	"github.com/goodaye/wire"
)

type S1 struct {
	wire.BaseService
}

func (s *S1) Init() error {
	fmt.Println("S1 Init ")
	return nil
}
func (s *S1) Start() error {
	fmt.Println("S1 Start ")

	return nil
}
func (s *S1) Stop() error {
	fmt.Println("S1 Stop ")
	return nil
}

type S2 struct {
	wire.BaseService
}

func (s *S2) Init() error {
	fmt.Println("S2 Init ")
	return nil
}
func (s *S2) Start() error {
	fmt.Println("S1 Start ")
	return nil
}
func (s *S2) Notify(msg wire.Message) error {

	fmt.Println("S2 notify:", msg)
	return nil
}

func init() {
	s1 := &S1{}
	s2 := &S2{}
	wire.Append(s1)
	wire.Append(s2)

}

func main() {
	fmt.Println("in main")
	go func() {
		fmt.Println("Start Run Service ")
		wire.Run()
	}()
	time.Sleep(1 * time.Second)
	msg := wire.Message{
		Key:  "NullMessage",
		Data: "change ",
	}
	wire.Notify(msg)
	wire.Stop()
}
