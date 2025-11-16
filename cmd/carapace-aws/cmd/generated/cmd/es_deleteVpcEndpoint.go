package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var es_deleteVpcEndpointCmd = &cobra.Command{
	Use:   "delete-vpc-endpoint",
	Short: "Deletes an Amazon OpenSearch Service-managed interface VPC endpoint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(es_deleteVpcEndpointCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(es_deleteVpcEndpointCmd).Standalone()

		es_deleteVpcEndpointCmd.Flags().String("vpc-endpoint-id", "", "The unique identifier of the endpoint to be deleted.")
		es_deleteVpcEndpointCmd.MarkFlagRequired("vpc-endpoint-id")
	})
	esCmd.AddCommand(es_deleteVpcEndpointCmd)
}
