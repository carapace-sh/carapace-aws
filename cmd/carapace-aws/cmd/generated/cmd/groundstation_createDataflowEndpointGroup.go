package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var groundstation_createDataflowEndpointGroupCmd = &cobra.Command{
	Use:   "create-dataflow-endpoint-group",
	Short: "Creates a `DataflowEndpoint` group containing the specified list of `DataflowEndpoint` objects.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(groundstation_createDataflowEndpointGroupCmd).Standalone()

	groundstation_createDataflowEndpointGroupCmd.Flags().String("contact-post-pass-duration-seconds", "", "Amount of time, in seconds, after a contact ends that the Ground Station Dataflow Endpoint Group will be in a `POSTPASS` state.")
	groundstation_createDataflowEndpointGroupCmd.Flags().String("contact-pre-pass-duration-seconds", "", "Amount of time, in seconds, before a contact starts that the Ground Station Dataflow Endpoint Group will be in a `PREPASS` state.")
	groundstation_createDataflowEndpointGroupCmd.Flags().String("endpoint-details", "", "Endpoint details of each endpoint in the dataflow endpoint group.")
	groundstation_createDataflowEndpointGroupCmd.Flags().String("tags", "", "Tags of a dataflow endpoint group.")
	groundstation_createDataflowEndpointGroupCmd.MarkFlagRequired("endpoint-details")
	groundstationCmd.AddCommand(groundstation_createDataflowEndpointGroupCmd)
}
