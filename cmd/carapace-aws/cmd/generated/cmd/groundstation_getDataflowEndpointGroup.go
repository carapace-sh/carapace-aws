package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var groundstation_getDataflowEndpointGroupCmd = &cobra.Command{
	Use:   "get-dataflow-endpoint-group",
	Short: "Returns the dataflow endpoint group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(groundstation_getDataflowEndpointGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(groundstation_getDataflowEndpointGroupCmd).Standalone()

		groundstation_getDataflowEndpointGroupCmd.Flags().String("dataflow-endpoint-group-id", "", "UUID of a dataflow endpoint group.")
		groundstation_getDataflowEndpointGroupCmd.MarkFlagRequired("dataflow-endpoint-group-id")
	})
	groundstationCmd.AddCommand(groundstation_getDataflowEndpointGroupCmd)
}
