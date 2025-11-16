package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_createVpcconnectionCmd = &cobra.Command{
	Use:   "create-vpcconnection",
	Short: "Creates a new VPC connection.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_createVpcconnectionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(quicksight_createVpcconnectionCmd).Standalone()

		quicksight_createVpcconnectionCmd.Flags().String("aws-account-id", "", "The Amazon Web Services account ID of the account where you want to create a new VPC connection.")
		quicksight_createVpcconnectionCmd.Flags().String("dns-resolvers", "", "A list of IP addresses of DNS resolver endpoints for the VPC connection.")
		quicksight_createVpcconnectionCmd.Flags().String("name", "", "The display name for the VPC connection.")
		quicksight_createVpcconnectionCmd.Flags().String("role-arn", "", "The IAM role to associate with the VPC connection.")
		quicksight_createVpcconnectionCmd.Flags().String("security-group-ids", "", "A list of security group IDs for the VPC connection.")
		quicksight_createVpcconnectionCmd.Flags().String("subnet-ids", "", "A list of subnet IDs for the VPC connection.")
		quicksight_createVpcconnectionCmd.Flags().String("tags", "", "A map of the key-value pairs for the resource tag or tags assigned to the VPC connection.")
		quicksight_createVpcconnectionCmd.Flags().String("vpcconnection-id", "", "The ID of the VPC connection that you're creating.")
		quicksight_createVpcconnectionCmd.MarkFlagRequired("aws-account-id")
		quicksight_createVpcconnectionCmd.MarkFlagRequired("name")
		quicksight_createVpcconnectionCmd.MarkFlagRequired("role-arn")
		quicksight_createVpcconnectionCmd.MarkFlagRequired("security-group-ids")
		quicksight_createVpcconnectionCmd.MarkFlagRequired("subnet-ids")
		quicksight_createVpcconnectionCmd.MarkFlagRequired("vpcconnection-id")
	})
	quicksightCmd.AddCommand(quicksight_createVpcconnectionCmd)
}
