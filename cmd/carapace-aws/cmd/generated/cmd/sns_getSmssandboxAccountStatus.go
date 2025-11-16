package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sns_getSmssandboxAccountStatusCmd = &cobra.Command{
	Use:   "get-smssandbox-account-status",
	Short: "Retrieves the SMS sandbox status for the calling Amazon Web Services account in the target Amazon Web Services Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sns_getSmssandboxAccountStatusCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sns_getSmssandboxAccountStatusCmd).Standalone()

	})
	snsCmd.AddCommand(sns_getSmssandboxAccountStatusCmd)
}
