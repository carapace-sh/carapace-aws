package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securityhub_disableSecurityHubV2Cmd = &cobra.Command{
	Use:   "disable-security-hub-v2",
	Short: "Disable the service for the current Amazon Web Services Region or specified Amazon Web Services Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityhub_disableSecurityHubV2Cmd).Standalone()

	securityhubCmd.AddCommand(securityhub_disableSecurityHubV2Cmd)
}
