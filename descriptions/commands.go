package descriptions

import (
	"errors"
	"github.com/minkimipt/contrail-introspect-cli/collection"
)

type command struct {
	// name of command line argument to invoce this functionality
	Name string
	// function that returns cli.Command to make interactive usage possible
	DescFn collection.DescCollection
	//description what this command does and how to use it
	Help string
}

var Commands = [...]command{
	command{DescFn: VrouterDropstats(), Name: "vrouter-dropstats", Help: "Show dropstats of vrouter"},
	command{DescFn: AgentCpu(), Name: "agent-cpu", Help: "Query the agent about its cpu utilization"},
	command{DescFn: AgentMemory(), Name: "agent-memory", Help: "Query the agent about its memory utilization"},
	command{DescFn: CtrlIfmap(), Name: "controller-ifmap", Help: "Query the ifmap through the controller"},
	command{DescFn: AgentPing(), Name: "agent-ping", Help: "Generate tcp ping in a vrf"},
	command{DescFn: Route(), Name: "agent-route", Help: "Show routes on agent"},
	command{DescFn: Interface(), Name: "agent-itf", Help: "Show interfaces on agent"},
	command{DescFn: Si(), Name: "agent-si", Help: "Show service instances on agent"},
	command{DescFn: Vrf(), Name: "agent-vrf", Help: "Show vrfs on agent "},
	command{DescFn: Peering(), Name: "agent-peering", Help: "Peering with controller on agent"},
	command{DescFn: Vn(), Name: "agent-vn", Help: "Show virtual networks on agent"},
	command{DescFn: Mpls(), Name: "agent-mpls", Help: "Show mpls on agent"},
	command{DescFn: RiSummary(), Name: "controller-ri", Help: "Show routing instances on controller"},
	command{DescFn: CtrlRoute(), Name: "controller-route", Help: "Show routes on controller"},
	command{DescFn: CtrlRouteSummary(), Name: "controller-route-summary", Help: "Show routes summary on controller"},
}

func GetCommand(name string) (collection.DescCollection, error) {
	for _, cmd := range Commands {
		if cmd.Name == name {
			return cmd.DescFn, nil
		}
	}
	e := errors.New("descriptions.GetCommand: cannot get " + name + " from Commands")
	return collection.DescCollection{}, e
}

func GetFields(name string) (fields []string, e error) {
	for _, cmd := range Commands {
		if cmd.Name == name {
			return cmd.DescFn.DescElt.LongDetail.GetFields(), nil
		}
	}
	return []string{}, errors.New("descriptions.GetFields: cannot get fields for command " + name)
}
