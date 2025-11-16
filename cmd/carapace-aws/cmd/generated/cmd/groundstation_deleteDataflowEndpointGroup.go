package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var groundstation_deleteDataflowEndpointGroupCmd = &cobra.Command{
	Use:   "delete-dataflow-endpoint-group",
	Short: "Deletes a dataflow endpoint group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(groundstation_deleteDataflowEndpointGroupCmd).Standalone()

	groundstation_deleteDataflowEndpointGroupCmd.Flags().String("dataflow-endpoint-group-id", "", "UUID of a dataflow endpoint group.")
	groundstation_deleteDataflowEndpointGroupCmd.MarkFlagRequired("dataflow-endpoint-group-id")
	groundstationCmd.AddCommand(groundstation_deleteDataflowEndpointGroupCmd)
}
