package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspector2_createCodeSecurityScanConfigurationCmd = &cobra.Command{
	Use:   "create-code-security-scan-configuration",
	Short: "Creates a scan configuration for code security scanning.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspector2_createCodeSecurityScanConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(inspector2_createCodeSecurityScanConfigurationCmd).Standalone()

		inspector2_createCodeSecurityScanConfigurationCmd.Flags().String("configuration", "", "The configuration settings for the code security scan.")
		inspector2_createCodeSecurityScanConfigurationCmd.Flags().String("level", "", "The security level for the scan configuration.")
		inspector2_createCodeSecurityScanConfigurationCmd.Flags().String("name", "", "The name of the scan configuration.")
		inspector2_createCodeSecurityScanConfigurationCmd.Flags().String("scope-settings", "", "The scope settings that define which repositories will be scanned.")
		inspector2_createCodeSecurityScanConfigurationCmd.Flags().String("tags", "", "The tags to apply to the scan configuration.")
		inspector2_createCodeSecurityScanConfigurationCmd.MarkFlagRequired("configuration")
		inspector2_createCodeSecurityScanConfigurationCmd.MarkFlagRequired("level")
		inspector2_createCodeSecurityScanConfigurationCmd.MarkFlagRequired("name")
	})
	inspector2Cmd.AddCommand(inspector2_createCodeSecurityScanConfigurationCmd)
}
