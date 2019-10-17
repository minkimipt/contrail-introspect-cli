package main

import "fmt"
import "log"
import "github.com/minkimipt/contrail-introspect-cli/descriptions"
import "github.com/minkimipt/contrail-introspect-cli/collection"

//import "reflect"

func main() {
	intfDescCollection := descriptions.XmppCount()
	//intfDescCollection := descriptions.Peering()
	var page collection.Sourcer
	args := []string{"172.25.151.84"}
	page = intfDescCollection.PageBuilder(args)
	col, e := page.Load(intfDescCollection)
	if e != nil {
		log.Fatal(e)
	}
	var list collection.Shower
	list = col
	vars := []string{}
	list.Long(80)
	list.Vars(&vars)
	fmt.Print(vars)
}
