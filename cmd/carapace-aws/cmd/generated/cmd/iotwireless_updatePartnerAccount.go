package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotwireless_updatePartnerAccountCmd = &cobra.Command{
	Use:   "update-partner-account",
	Short: "Updates properties of a partner account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotwireless_updatePartnerAccountCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotwireless_updatePartnerAccountCmd).Standalone()

		iotwireless_updatePartnerAccountCmd.Flags().String("partner-account-id", "", "The ID of the partner account to update.")
		iotwireless_updatePartnerAccountCmd.Flags().String("partner-type", "", "The partner type.")
		iotwireless_updatePartnerAccountCmd.Flags().String("sidewalk", "", "The Sidewalk account credentials.")
		iotwireless_updatePartnerAccountCmd.MarkFlagRequired("partner-account-id")
		iotwireless_updatePartnerAccountCmd.MarkFlagRequired("partner-type")
		iotwireless_updatePartnerAccountCmd.MarkFlagRequired("sidewalk")
	})
	iotwirelessCmd.AddCommand(iotwireless_updatePartnerAccountCmd)
}
