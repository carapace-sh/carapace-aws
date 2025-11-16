package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var opensearch_deleteVpcEndpointCmd = &cobra.Command{
	Use:   "delete-vpc-endpoint",
	Short: "Deletes an Amazon OpenSearch Service-managed interface VPC endpoint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opensearch_deleteVpcEndpointCmd).Standalone()

	opensearch_deleteVpcEndpointCmd.Flags().String("vpc-endpoint-id", "", "The unique identifier of the endpoint.")
	opensearch_deleteVpcEndpointCmd.MarkFlagRequired("vpc-endpoint-id")
	opensearchCmd.AddCommand(opensearch_deleteVpcEndpointCmd)
}
