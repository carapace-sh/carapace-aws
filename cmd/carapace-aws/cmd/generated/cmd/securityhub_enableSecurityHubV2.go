package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securityhub_enableSecurityHubV2Cmd = &cobra.Command{
	Use:   "enable-security-hub-v2",
	Short: "Enables the service in account for the current Amazon Web Services Region or specified Amazon Web Services Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityhub_enableSecurityHubV2Cmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(securityhub_enableSecurityHubV2Cmd).Standalone()

		securityhub_enableSecurityHubV2Cmd.Flags().String("tags", "", "The tags to add to the hub V2 resource when you enable Security Hub.")
	})
	securityhubCmd.AddCommand(securityhub_enableSecurityHubV2Cmd)
}
