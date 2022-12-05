package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Action = func(c *cli.Context) {
		if c.NArg() != 0 {
			fmt.Printf("未找到命令: %s\n运行命令 %s help 获取帮助\n", c.Args().Get(0), app.Name)
			return
		}

		var prompt string

		prompt = app.Name + " > "
	L:
		for {
			var input string
			fmt.Print(prompt)
			//   fmt.Scanln(&input)

			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan() // use `for scanner.Scan()` to keep reading
			input = scanner.Text()
			//fmt.Println("captured:",input)
			switch input {
			case "close":
				fmt.Println("close.")
				break L
			default:
			}
			//fmt.Print(input)
			cmdArgs := strings.Split(input, " ")
			//fmt.Print(len(cmdArgs))
			if len(cmdArgs) == 0 {
				continue
			}

			s := []string{app.Name}
			s = append(s, cmdArgs...)

			c.App.Run(s)

		}

		return
	}

	app.Commands = []cli.Command{
		{
			Name:    "add",
			Aliases: []string{"a"},
			Usage:   "add a task to the list",
			Action: func(c *cli.Context) error {
				fmt.Println("added task: ", c.Args().First())
				return nil
			},
		},
		{
			Name:    "complete",
			Aliases: []string{"c"},
			Usage:   "complete a task on the list",
			Action: func(c *cli.Context) error {
				fmt.Println("completed task: ", c.Args().First())
				return nil
			},
		},
		{
			Name:    "template",
			Aliases: []string{"t"},
			Usage:   "options for task templates",
			Subcommands: []cli.Command{
				{
					Name:  "add",
					Usage: "add a new template",
					Action: func(c *cli.Context) error {
						fmt.Println("new task template: ", c.Args().First())
						return nil
					},
				},
				{
					Name:  "remove",
					Usage: "remove an existing template",
					Action: func(c *cli.Context) error {
						fmt.Println("removed task template: ", c.Args().First())
						return nil
					},
				},
			},
		},
	}

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}

//func main() {
//	// Create a cache with a default expiration time of 5 minutes, and which
//	// purges expired items every 10 minutes
//	c := cache.New(5*time.Minute, 10*time.Minute)
//
//	// Set the value of the key "foo" to "bar", with the default expiration time
//	c.Set("foo", "bar", cache.DefaultExpiration)
//
//	// Set the value of the key "baz" to 42, with no expiration time
//	// (the item won't be removed until it is re-set, or removed using
//	// c.Delete("baz")
//	c.Set("baz", 42, cache.NoExpiration)
//
//	// Get the string associated with the key "foo" from the cache
//	foo, found := c.Get("foo")
//	if found {
//		fmt.Println(foo)
//	}
//
//	// Since Go is statically typed, and cache values can be anything, type
//	// assertion is needed when values are being passed to functions that don't
//	// take arbitrary types, (i.e. interface{}). The simplest way to do this for
//	// values which will only be used once--e.g. for passing to another
//	// function--is:
//	foo, found := c.Get("foo")
//	if found {
//		MyFunction(foo.(string))
//	}
//
//	// This gets tedious if the value is used several times in the same function.
//	// You might do either of the following instead:
//	if x, found := c.Get("foo"); found {
//		foo := x.(string)
//		// ...
//	}
//	// or
//	var foo string
//	if x, found := c.Get("foo"); found {
//		foo = x.(string)
//	}
//	// ...
//	// foo can then be passed around freely as a string
//
//	// Want performance? Store pointers!
//	c.Set("foo", &MyStruct, cache.DefaultExpiration)
//	if x, found := c.Get("foo"); found {
//		foo := x.(*MyStruct)
//		// ...
//	}
//}
