package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_deleteVpcconnectionCmd = &cobra.Command{
	Use:   "delete-vpcconnection",
	Short: "Deletes a VPC connection.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_deleteVpcconnectionCmd).Standalone()

	quicksight_deleteVpcconnectionCmd.Flags().String("aws-account-id", "", "The Amazon Web Services account ID of the account where you want to delete a VPC connection.")
	quicksight_deleteVpcconnectionCmd.Flags().String("vpcconnection-id", "", "The ID of the VPC connection that you're creating.")
	quicksight_deleteVpcconnectionCmd.MarkFlagRequired("aws-account-id")
	quicksight_deleteVpcconnectionCmd.MarkFlagRequired("vpcconnection-id")
	quicksightCmd.AddCommand(quicksight_deleteVpcconnectionCmd)
}
