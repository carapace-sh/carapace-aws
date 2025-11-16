package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspector2_getCodeSecurityScanConfigurationCmd = &cobra.Command{
	Use:   "get-code-security-scan-configuration",
	Short: "Retrieves information about a code security scan configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspector2_getCodeSecurityScanConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(inspector2_getCodeSecurityScanConfigurationCmd).Standalone()

		inspector2_getCodeSecurityScanConfigurationCmd.Flags().String("scan-configuration-arn", "", "The Amazon Resource Name (ARN) of the scan configuration to retrieve.")
		inspector2_getCodeSecurityScanConfigurationCmd.MarkFlagRequired("scan-configuration-arn")
	})
	inspector2Cmd.AddCommand(inspector2_getCodeSecurityScanConfigurationCmd)
}
