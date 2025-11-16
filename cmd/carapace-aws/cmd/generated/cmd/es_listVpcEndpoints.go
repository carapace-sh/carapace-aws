package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var es_listVpcEndpointsCmd = &cobra.Command{
	Use:   "list-vpc-endpoints",
	Short: "Retrieves all Amazon OpenSearch Service-managed VPC endpoints in the current account and Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(es_listVpcEndpointsCmd).Standalone()

	es_listVpcEndpointsCmd.Flags().String("next-token", "", "Identifier to allow retrieval of paginated results.")
	esCmd.AddCommand(es_listVpcEndpointsCmd)
}
