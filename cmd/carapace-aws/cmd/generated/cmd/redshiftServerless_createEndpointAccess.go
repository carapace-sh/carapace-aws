package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshiftServerless_createEndpointAccessCmd = &cobra.Command{
	Use:   "create-endpoint-access",
	Short: "Creates an Amazon Redshift Serverless managed VPC endpoint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshiftServerless_createEndpointAccessCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(redshiftServerless_createEndpointAccessCmd).Standalone()

		redshiftServerless_createEndpointAccessCmd.Flags().String("endpoint-name", "", "The name of the VPC endpoint.")
		redshiftServerless_createEndpointAccessCmd.Flags().String("owner-account", "", "The owner Amazon Web Services account for the Amazon Redshift Serverless workgroup.")
		redshiftServerless_createEndpointAccessCmd.Flags().String("subnet-ids", "", "The unique identifers of subnets from which Amazon Redshift Serverless chooses one to deploy a VPC endpoint.")
		redshiftServerless_createEndpointAccessCmd.Flags().String("vpc-security-group-ids", "", "The unique identifiers of the security group that defines the ports, protocols, and sources for inbound traffic that you are authorizing into your endpoint.")
		redshiftServerless_createEndpointAccessCmd.Flags().String("workgroup-name", "", "The name of the workgroup to associate with the VPC endpoint.")
		redshiftServerless_createEndpointAccessCmd.MarkFlagRequired("endpoint-name")
		redshiftServerless_createEndpointAccessCmd.MarkFlagRequired("subnet-ids")
		redshiftServerless_createEndpointAccessCmd.MarkFlagRequired("workgroup-name")
	})
	redshiftServerlessCmd.AddCommand(redshiftServerless_createEndpointAccessCmd)
}
