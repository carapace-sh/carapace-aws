package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointEmail_getAccountCmd = &cobra.Command{
	Use:   "get-account",
	Short: "Obtain information about the email-sending status and capabilities of your Amazon Pinpoint account in the current AWS Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointEmail_getAccountCmd).Standalone()

	pinpointEmailCmd.AddCommand(pinpointEmail_getAccountCmd)
}
