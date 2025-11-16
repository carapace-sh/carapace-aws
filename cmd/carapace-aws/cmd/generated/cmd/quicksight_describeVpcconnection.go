package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_describeVpcconnectionCmd = &cobra.Command{
	Use:   "describe-vpcconnection",
	Short: "Describes a VPC connection.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_describeVpcconnectionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(quicksight_describeVpcconnectionCmd).Standalone()

		quicksight_describeVpcconnectionCmd.Flags().String("aws-account-id", "", "The Amazon Web Services account ID of the account that contains the VPC connection that you want described.")
		quicksight_describeVpcconnectionCmd.Flags().String("vpcconnection-id", "", "The ID of the VPC connection that you're creating.")
		quicksight_describeVpcconnectionCmd.MarkFlagRequired("aws-account-id")
		quicksight_describeVpcconnectionCmd.MarkFlagRequired("vpcconnection-id")
	})
	quicksightCmd.AddCommand(quicksight_describeVpcconnectionCmd)
}
