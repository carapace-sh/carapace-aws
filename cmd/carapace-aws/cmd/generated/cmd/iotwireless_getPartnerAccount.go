package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotwireless_getPartnerAccountCmd = &cobra.Command{
	Use:   "get-partner-account",
	Short: "Gets information about a partner account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotwireless_getPartnerAccountCmd).Standalone()

	iotwireless_getPartnerAccountCmd.Flags().String("partner-account-id", "", "The partner account ID to disassociate from the AWS account.")
	iotwireless_getPartnerAccountCmd.Flags().String("partner-type", "", "The partner type.")
	iotwireless_getPartnerAccountCmd.MarkFlagRequired("partner-account-id")
	iotwireless_getPartnerAccountCmd.MarkFlagRequired("partner-type")
	iotwirelessCmd.AddCommand(iotwireless_getPartnerAccountCmd)
}
