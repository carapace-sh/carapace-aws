package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var opensearchserverless_deleteVpcEndpointCmd = &cobra.Command{
	Use:   "delete-vpc-endpoint",
	Short: "Deletes an OpenSearch Serverless-managed interface endpoint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opensearchserverless_deleteVpcEndpointCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(opensearchserverless_deleteVpcEndpointCmd).Standalone()

		opensearchserverless_deleteVpcEndpointCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier to ensure idempotency of the request.")
		opensearchserverless_deleteVpcEndpointCmd.Flags().String("id", "", "The VPC endpoint identifier.")
		opensearchserverless_deleteVpcEndpointCmd.MarkFlagRequired("id")
	})
	opensearchserverlessCmd.AddCommand(opensearchserverless_deleteVpcEndpointCmd)
}
