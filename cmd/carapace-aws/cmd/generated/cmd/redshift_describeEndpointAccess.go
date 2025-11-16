package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_describeEndpointAccessCmd = &cobra.Command{
	Use:   "describe-endpoint-access",
	Short: "Describes a Redshift-managed VPC endpoint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_describeEndpointAccessCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(redshift_describeEndpointAccessCmd).Standalone()

		redshift_describeEndpointAccessCmd.Flags().String("cluster-identifier", "", "The cluster identifier associated with the described endpoint.")
		redshift_describeEndpointAccessCmd.Flags().String("endpoint-name", "", "The name of the endpoint to be described.")
		redshift_describeEndpointAccessCmd.Flags().String("marker", "", "An optional pagination token provided by a previous `DescribeEndpointAccess` request.")
		redshift_describeEndpointAccessCmd.Flags().String("max-records", "", "The maximum number of records to include in the response.")
		redshift_describeEndpointAccessCmd.Flags().String("resource-owner", "", "The Amazon Web Services account ID of the owner of the cluster.")
		redshift_describeEndpointAccessCmd.Flags().String("vpc-id", "", "The virtual private cloud (VPC) identifier with access to the cluster.")
	})
	redshiftCmd.AddCommand(redshift_describeEndpointAccessCmd)
}
