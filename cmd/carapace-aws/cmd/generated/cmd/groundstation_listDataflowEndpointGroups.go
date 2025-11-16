package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var groundstation_listDataflowEndpointGroupsCmd = &cobra.Command{
	Use:   "list-dataflow-endpoint-groups",
	Short: "Returns a list of `DataflowEndpoint` groups.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(groundstation_listDataflowEndpointGroupsCmd).Standalone()

	groundstation_listDataflowEndpointGroupsCmd.Flags().String("max-results", "", "Maximum number of dataflow endpoint groups returned.")
	groundstation_listDataflowEndpointGroupsCmd.Flags().String("next-token", "", "Next token returned in the request of a previous `ListDataflowEndpointGroups` call.")
	groundstationCmd.AddCommand(groundstation_listDataflowEndpointGroupsCmd)
}
