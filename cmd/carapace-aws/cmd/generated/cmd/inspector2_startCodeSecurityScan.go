package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspector2_startCodeSecurityScanCmd = &cobra.Command{
	Use:   "start-code-security-scan",
	Short: "Initiates a code security scan on a specified repository.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspector2_startCodeSecurityScanCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(inspector2_startCodeSecurityScanCmd).Standalone()

		inspector2_startCodeSecurityScanCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		inspector2_startCodeSecurityScanCmd.Flags().String("resource", "", "The resource identifier for the code repository to scan.")
		inspector2_startCodeSecurityScanCmd.MarkFlagRequired("resource")
	})
	inspector2Cmd.AddCommand(inspector2_startCodeSecurityScanCmd)
}
