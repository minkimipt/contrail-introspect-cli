package main

import "log"
import "github.com/nlewo/contrail-introspect-cli/descriptions"
import "github.com/nlewo/contrail-introspect-cli/collection"

func main() {
	//	intfDescCollection := descriptions.AgentCpu()
	intfDescCollection := descriptions.AgentCpu()
	var page collection.Sourcer
	args := []string{"172.25.151.84"}
	page = intfDescCollection.PageBuilder(args)
	col, e := page.Load(intfDescCollection)
	if e != nil {
		log.Fatal(e)
	}
	var list collection.Shower
	list = col
	list.Short()
	list.Long(80)
}
