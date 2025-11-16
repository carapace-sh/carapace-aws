package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspector2_batchAssociateCodeSecurityScanConfigurationCmd = &cobra.Command{
	Use:   "batch-associate-code-security-scan-configuration",
	Short: "Associates multiple code repositories with an Amazon Inspector code security scan configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspector2_batchAssociateCodeSecurityScanConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(inspector2_batchAssociateCodeSecurityScanConfigurationCmd).Standalone()

		inspector2_batchAssociateCodeSecurityScanConfigurationCmd.Flags().String("associate-configuration-requests", "", "A list of code repositories to associate with the specified scan configuration.")
		inspector2_batchAssociateCodeSecurityScanConfigurationCmd.MarkFlagRequired("associate-configuration-requests")
	})
	inspector2Cmd.AddCommand(inspector2_batchAssociateCodeSecurityScanConfigurationCmd)
}
