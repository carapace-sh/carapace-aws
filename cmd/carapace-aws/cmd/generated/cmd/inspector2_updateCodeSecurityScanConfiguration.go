package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspector2_updateCodeSecurityScanConfigurationCmd = &cobra.Command{
	Use:   "update-code-security-scan-configuration",
	Short: "Updates an existing code security scan configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspector2_updateCodeSecurityScanConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(inspector2_updateCodeSecurityScanConfigurationCmd).Standalone()

		inspector2_updateCodeSecurityScanConfigurationCmd.Flags().String("configuration", "", "The updated configuration settings for the code security scan.")
		inspector2_updateCodeSecurityScanConfigurationCmd.Flags().String("scan-configuration-arn", "", "The Amazon Resource Name (ARN) of the scan configuration to update.")
		inspector2_updateCodeSecurityScanConfigurationCmd.MarkFlagRequired("configuration")
		inspector2_updateCodeSecurityScanConfigurationCmd.MarkFlagRequired("scan-configuration-arn")
	})
	inspector2Cmd.AddCommand(inspector2_updateCodeSecurityScanConfigurationCmd)
}
