package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var groundstation_createDataflowEndpointGroupV2Cmd = &cobra.Command{
	Use:   "create-dataflow-endpoint-group-v2",
	Short: "Creates a `DataflowEndpointGroupV2` containing the specified list of `DataflowEndpoint` objects.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(groundstation_createDataflowEndpointGroupV2Cmd).Standalone()

	groundstation_createDataflowEndpointGroupV2Cmd.Flags().String("contact-post-pass-duration-seconds", "", "Amount of time, in seconds, after a contact ends that the Ground Station Dataflow Endpoint Group will be in a `POSTPASS` state.")
	groundstation_createDataflowEndpointGroupV2Cmd.Flags().String("contact-pre-pass-duration-seconds", "", "Amount of time, in seconds, before a contact starts that the Ground Station Dataflow Endpoint Group will be in a `PREPASS` state.")
	groundstation_createDataflowEndpointGroupV2Cmd.Flags().String("endpoints", "", "Dataflow endpoint group's endpoint definitions")
	groundstation_createDataflowEndpointGroupV2Cmd.Flags().String("tags", "", "Tags of a V2 dataflow endpoint group.")
	groundstation_createDataflowEndpointGroupV2Cmd.MarkFlagRequired("endpoints")
	groundstationCmd.AddCommand(groundstation_createDataflowEndpointGroupV2Cmd)
}
