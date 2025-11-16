package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var osis_listPipelineEndpointConnectionsCmd = &cobra.Command{
	Use:   "list-pipeline-endpoint-connections",
	Short: "Lists the pipeline endpoints connected to pipelines in your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(osis_listPipelineEndpointConnectionsCmd).Standalone()

	osis_listPipelineEndpointConnectionsCmd.Flags().String("max-results", "", "The maximum number of pipeline endpoint connections to return in the response.")
	osis_listPipelineEndpointConnectionsCmd.Flags().String("next-token", "", "If your initial `ListPipelineEndpointConnections` operation returns a `nextToken`, you can include the returned `nextToken` in subsequent `ListPipelineEndpointConnections` operations, which returns results in the next page.")
	osisCmd.AddCommand(osis_listPipelineEndpointConnectionsCmd)
}
