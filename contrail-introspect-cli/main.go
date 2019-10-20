package main

import "fmt"
import "os"
import "log"

import "github.com/jbowtie/gokogiri/xpath"
import cli "gopkg.in/urfave/cli.v2"

import "github.com/minkimipt/contrail-introspect-cli/descriptions"
import "github.com/minkimipt/contrail-introspect-cli/utils"
import "github.com/minkimipt/contrail-introspect-cli/collection"

func multiple(vrouter string, vrf_name string, count bool) {
	url := "http://" + vrouter + ":8085" + "/Snh_PageReq?x=begin:-1,end:-1,table:" + vrf_name + ".uc.route.0,"

	doc, err := collection.Load(url, false)
	if err != nil {
		fmt.Printf("%s\n", err)
	} else {
		defer doc.Free()
		xps := xpath.Compile("//route_list/list/RouteUcSandeshData/path_list/list/PathSandeshData/nh/NhSandeshData/mc_list/../../../../../../src_ip/text()")
		ss, _ := doc.Root().Search(xps)
		if count {
			fmt.Printf("%d\n", len(ss))
		} else {
			for _, s := range ss {
				fmt.Printf("%s\n", s)
			}
		}
	}
}

func main() {
	var count bool
	var hosts_file string

	app := (&cli.App{})
	app.Name = "contrail-introspect-cli"
	app.Usage = "CLI on ContraiL Introspects"
	app.Version = "0.0.4"
	app.EnableShellCompletion = true
	app.Before = func(c *cli.Context) error {
		if c.IsSet("hosts") {
			var err error
			utils.HostMap, err = utils.LoadHostsFile(c.String("hosts"))
			return err
		}
		return nil
	}
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:        "hosts",
			Usage:       "host file to do DNS resolution",
			Destination: &hosts_file,
		}}
	app.Commands = []*cli.Command{}
	for _, command := range descriptions.Commands {
		app.Commands = append(app.Commands, GenCommand(command.DescFn, command.Name, command.Help))
	}
	Commands_other := []*cli.Command{
		descriptions.Ping(),
		descriptions.Follow(),
		descriptions.Path(),
		{
			Name:      "agent-multiple",
			Usage:     "List routes with multiple nexthops",
			ArgsUsage: "vrouter vrf_name",
			Flags: []cli.Flag{
				&cli.BoolFlag{
					Name:        "count",
					Destination: &count,
				}},
			Action: func(c *cli.Context) error {
				if c.NArg() != 2 {
					log.Fatal("Wrong argument number!")
				}
				vrouter := c.Args().Get(0)
				vrf_name := c.Args().Get(1)
				multiple(vrouter, vrf_name, count)
				return nil
			},
		},
		descriptions.RouteDiff(),
	}
	app.Commands = append(app.Commands, Commands_other...)
	app.Run(os.Args)
}
