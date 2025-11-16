package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_updateVpcconnectionCmd = &cobra.Command{
	Use:   "update-vpcconnection",
	Short: "Updates a VPC connection.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_updateVpcconnectionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(quicksight_updateVpcconnectionCmd).Standalone()

		quicksight_updateVpcconnectionCmd.Flags().String("aws-account-id", "", "The Amazon Web Services account ID of the account that contains the VPC connection that you want to update.")
		quicksight_updateVpcconnectionCmd.Flags().String("dns-resolvers", "", "A list of IP addresses of DNS resolver endpoints for the VPC connection.")
		quicksight_updateVpcconnectionCmd.Flags().String("name", "", "The display name for the VPC connection.")
		quicksight_updateVpcconnectionCmd.Flags().String("role-arn", "", "An IAM role associated with the VPC connection.")
		quicksight_updateVpcconnectionCmd.Flags().String("security-group-ids", "", "A list of security group IDs for the VPC connection.")
		quicksight_updateVpcconnectionCmd.Flags().String("subnet-ids", "", "A list of subnet IDs for the VPC connection.")
		quicksight_updateVpcconnectionCmd.Flags().String("vpcconnection-id", "", "The ID of the VPC connection that you're updating.")
		quicksight_updateVpcconnectionCmd.MarkFlagRequired("aws-account-id")
		quicksight_updateVpcconnectionCmd.MarkFlagRequired("name")
		quicksight_updateVpcconnectionCmd.MarkFlagRequired("role-arn")
		quicksight_updateVpcconnectionCmd.MarkFlagRequired("security-group-ids")
		quicksight_updateVpcconnectionCmd.MarkFlagRequired("subnet-ids")
		quicksight_updateVpcconnectionCmd.MarkFlagRequired("vpcconnection-id")
	})
	quicksightCmd.AddCommand(quicksight_updateVpcconnectionCmd)
}
