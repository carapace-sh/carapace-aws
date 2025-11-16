package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securityhub_updateSecurityHubConfigurationCmd = &cobra.Command{
	Use:   "update-security-hub-configuration",
	Short: "Updates configuration options for Security Hub.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityhub_updateSecurityHubConfigurationCmd).Standalone()

	securityhub_updateSecurityHubConfigurationCmd.Flags().Bool("auto-enable-controls", false, "Whether to automatically enable new controls when they are added to standards that are enabled.")
	securityhub_updateSecurityHubConfigurationCmd.Flags().String("control-finding-generator", "", "Updates whether the calling account has consolidated control findings turned on.")
	securityhub_updateSecurityHubConfigurationCmd.Flags().Bool("no-auto-enable-controls", false, "Whether to automatically enable new controls when they are added to standards that are enabled.")
	securityhub_updateSecurityHubConfigurationCmd.Flag("no-auto-enable-controls").Hidden = true
	securityhubCmd.AddCommand(securityhub_updateSecurityHubConfigurationCmd)
}
