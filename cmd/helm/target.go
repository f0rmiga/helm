package main

import (
	"fmt"

	"github.com/deis/helm-dm/pkg/format"
	"github.com/deis/helm-dm/pkg/kubectl"
)

func target(dryRun bool) error {
	client := kubectl.Client
	if dryRun {
		client = kubectl.PrintRunner{}
	}
	out, err := client.ClusterInfo()
	if err != nil {
		return fmt.Errorf("%s (%s)", out, err)
	}
	format.Msg(string(out))
	return nil
}
