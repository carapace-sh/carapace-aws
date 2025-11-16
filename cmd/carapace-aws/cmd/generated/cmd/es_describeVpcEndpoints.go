package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var es_describeVpcEndpointsCmd = &cobra.Command{
	Use:   "describe-vpc-endpoints",
	Short: "Describes one or more Amazon OpenSearch Service-managed VPC endpoints.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(es_describeVpcEndpointsCmd).Standalone()

	es_describeVpcEndpointsCmd.Flags().String("vpc-endpoint-ids", "", "The unique identifiers of the endpoints to get information about.")
	es_describeVpcEndpointsCmd.MarkFlagRequired("vpc-endpoint-ids")
	esCmd.AddCommand(es_describeVpcEndpointsCmd)
}
