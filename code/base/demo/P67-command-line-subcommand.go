package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	// 声明一个命令子集合
	fooCmd := flag.NewFlagSet("foo", flag.ExitOnError)
	// 在该子命令下声明flag
	fooEnable := fooCmd.Bool("enable",false,"enable")
	fooName := fooCmd.String("name","","name")

	barCmd := flag.NewFlagSet("bar",flag.ExitOnError)
	barLevel := barCmd.Int("level",0,"level")

	// 必须要有子命令
	if len(os.Args) < 2 {
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "foo":
		fooCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'foo'")
		fmt.Println(" enable:",*fooEnable)
		fmt.Println(" name:",*fooName)
		fmt.Println(" tail:", fooCmd.Args())
	case "bar":
		barCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'bar'")
		fmt.Println(" level:",*barLevel)
		fmt.Println(" tail:",barCmd.Args())
	default:
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}

}
