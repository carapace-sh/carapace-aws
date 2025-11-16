package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securityhub_disableSecurityHubCmd = &cobra.Command{
	Use:   "disable-security-hub",
	Short: "Disables Security Hub in your account only in the current Amazon Web Services Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityhub_disableSecurityHubCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(securityhub_disableSecurityHubCmd).Standalone()

	})
	securityhubCmd.AddCommand(securityhub_disableSecurityHubCmd)
}
