package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var opensearchserverless_batchGetVpcEndpointCmd = &cobra.Command{
	Use:   "batch-get-vpc-endpoint",
	Short: "Returns attributes for one or more VPC endpoints associated with the current account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opensearchserverless_batchGetVpcEndpointCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(opensearchserverless_batchGetVpcEndpointCmd).Standalone()

		opensearchserverless_batchGetVpcEndpointCmd.Flags().String("ids", "", "A list of VPC endpoint identifiers.")
		opensearchserverless_batchGetVpcEndpointCmd.MarkFlagRequired("ids")
	})
	opensearchserverlessCmd.AddCommand(opensearchserverless_batchGetVpcEndpointCmd)
}
