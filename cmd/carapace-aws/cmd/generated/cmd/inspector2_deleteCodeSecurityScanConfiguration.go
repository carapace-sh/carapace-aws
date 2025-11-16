package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspector2_deleteCodeSecurityScanConfigurationCmd = &cobra.Command{
	Use:   "delete-code-security-scan-configuration",
	Short: "Deletes a code security scan configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspector2_deleteCodeSecurityScanConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(inspector2_deleteCodeSecurityScanConfigurationCmd).Standalone()

		inspector2_deleteCodeSecurityScanConfigurationCmd.Flags().String("scan-configuration-arn", "", "The Amazon Resource Name (ARN) of the scan configuration to delete.")
		inspector2_deleteCodeSecurityScanConfigurationCmd.MarkFlagRequired("scan-configuration-arn")
	})
	inspector2Cmd.AddCommand(inspector2_deleteCodeSecurityScanConfigurationCmd)
}
