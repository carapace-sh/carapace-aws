package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var opensearch_describeVpcEndpointsCmd = &cobra.Command{
	Use:   "describe-vpc-endpoints",
	Short: "Describes one or more Amazon OpenSearch Service-managed VPC endpoints.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opensearch_describeVpcEndpointsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(opensearch_describeVpcEndpointsCmd).Standalone()

		opensearch_describeVpcEndpointsCmd.Flags().String("vpc-endpoint-ids", "", "The unique identifiers of the endpoints to get information about.")
		opensearch_describeVpcEndpointsCmd.MarkFlagRequired("vpc-endpoint-ids")
	})
	opensearchCmd.AddCommand(opensearch_describeVpcEndpointsCmd)
}
