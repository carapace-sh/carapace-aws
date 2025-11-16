package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_createEndpointAccessCmd = &cobra.Command{
	Use:   "create-endpoint-access",
	Short: "Creates a Redshift-managed VPC endpoint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_createEndpointAccessCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(redshift_createEndpointAccessCmd).Standalone()

		redshift_createEndpointAccessCmd.Flags().String("cluster-identifier", "", "The cluster identifier of the cluster to access.")
		redshift_createEndpointAccessCmd.Flags().String("endpoint-name", "", "The Redshift-managed VPC endpoint name.")
		redshift_createEndpointAccessCmd.Flags().String("resource-owner", "", "The Amazon Web Services account ID of the owner of the cluster.")
		redshift_createEndpointAccessCmd.Flags().String("subnet-group-name", "", "The subnet group from which Amazon Redshift chooses the subnet to deploy the endpoint.")
		redshift_createEndpointAccessCmd.Flags().String("vpc-security-group-ids", "", "The security group that defines the ports, protocols, and sources for inbound traffic that you are authorizing into your endpoint.")
		redshift_createEndpointAccessCmd.MarkFlagRequired("endpoint-name")
		redshift_createEndpointAccessCmd.MarkFlagRequired("subnet-group-name")
	})
	redshiftCmd.AddCommand(redshift_createEndpointAccessCmd)
}
