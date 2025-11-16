package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotwireless_disassociateAwsAccountFromPartnerAccountCmd = &cobra.Command{
	Use:   "disassociate-aws-account-from-partner-account",
	Short: "Disassociates your AWS account from a partner account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotwireless_disassociateAwsAccountFromPartnerAccountCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotwireless_disassociateAwsAccountFromPartnerAccountCmd).Standalone()

		iotwireless_disassociateAwsAccountFromPartnerAccountCmd.Flags().String("partner-account-id", "", "The partner account ID to disassociate from the AWS account.")
		iotwireless_disassociateAwsAccountFromPartnerAccountCmd.Flags().String("partner-type", "", "The partner type.")
		iotwireless_disassociateAwsAccountFromPartnerAccountCmd.MarkFlagRequired("partner-account-id")
		iotwireless_disassociateAwsAccountFromPartnerAccountCmd.MarkFlagRequired("partner-type")
	})
	iotwirelessCmd.AddCommand(iotwireless_disassociateAwsAccountFromPartnerAccountCmd)
}
