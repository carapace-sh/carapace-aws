package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sesv2_putAccountSendingAttributesCmd = &cobra.Command{
	Use:   "put-account-sending-attributes",
	Short: "Enable or disable the ability of your account to send email.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sesv2_putAccountSendingAttributesCmd).Standalone()

	sesv2_putAccountSendingAttributesCmd.Flags().String("sending-enabled", "", "Enables or disables your account's ability to send email.")
	sesv2Cmd.AddCommand(sesv2_putAccountSendingAttributesCmd)
}
