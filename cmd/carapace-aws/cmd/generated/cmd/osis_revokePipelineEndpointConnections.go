package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var osis_revokePipelineEndpointConnectionsCmd = &cobra.Command{
	Use:   "revoke-pipeline-endpoint-connections",
	Short: "Revokes pipeline endpoints from specified endpoint IDs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(osis_revokePipelineEndpointConnectionsCmd).Standalone()

	osis_revokePipelineEndpointConnectionsCmd.Flags().String("endpoint-ids", "", "A list of endpoint IDs for which to revoke access to the pipeline.")
	osis_revokePipelineEndpointConnectionsCmd.Flags().String("pipeline-arn", "", "The Amazon Resource Name (ARN) of the pipeline from which to revoke endpoint connections.")
	osis_revokePipelineEndpointConnectionsCmd.MarkFlagRequired("endpoint-ids")
	osis_revokePipelineEndpointConnectionsCmd.MarkFlagRequired("pipeline-arn")
	osisCmd.AddCommand(osis_revokePipelineEndpointConnectionsCmd)
}
