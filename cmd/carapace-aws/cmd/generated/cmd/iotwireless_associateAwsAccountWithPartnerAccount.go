package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotwireless_associateAwsAccountWithPartnerAccountCmd = &cobra.Command{
	Use:   "associate-aws-account-with-partner-account",
	Short: "Associates a partner account with your AWS account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotwireless_associateAwsAccountWithPartnerAccountCmd).Standalone()

	iotwireless_associateAwsAccountWithPartnerAccountCmd.Flags().String("client-request-token", "", "Each resource must have a unique client request token.")
	iotwireless_associateAwsAccountWithPartnerAccountCmd.Flags().String("sidewalk", "", "The Sidewalk account credentials.")
	iotwireless_associateAwsAccountWithPartnerAccountCmd.Flags().String("tags", "", "The tags to attach to the specified resource.")
	iotwireless_associateAwsAccountWithPartnerAccountCmd.MarkFlagRequired("sidewalk")
	iotwirelessCmd.AddCommand(iotwireless_associateAwsAccountWithPartnerAccountCmd)
}
