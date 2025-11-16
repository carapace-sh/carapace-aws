package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspector2_batchDisassociateCodeSecurityScanConfigurationCmd = &cobra.Command{
	Use:   "batch-disassociate-code-security-scan-configuration",
	Short: "Disassociates multiple code repositories from an Amazon Inspector code security scan configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspector2_batchDisassociateCodeSecurityScanConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(inspector2_batchDisassociateCodeSecurityScanConfigurationCmd).Standalone()

		inspector2_batchDisassociateCodeSecurityScanConfigurationCmd.Flags().String("disassociate-configuration-requests", "", "A list of code repositories to disassociate from the specified scan configuration.")
		inspector2_batchDisassociateCodeSecurityScanConfigurationCmd.MarkFlagRequired("disassociate-configuration-requests")
	})
	inspector2Cmd.AddCommand(inspector2_batchDisassociateCodeSecurityScanConfigurationCmd)
}
