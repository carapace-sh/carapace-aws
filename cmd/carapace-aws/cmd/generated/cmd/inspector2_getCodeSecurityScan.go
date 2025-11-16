package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspector2_getCodeSecurityScanCmd = &cobra.Command{
	Use:   "get-code-security-scan",
	Short: "Retrieves information about a specific code security scan.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspector2_getCodeSecurityScanCmd).Standalone()

	inspector2_getCodeSecurityScanCmd.Flags().String("resource", "", "The resource identifier for the code repository that was scanned.")
	inspector2_getCodeSecurityScanCmd.Flags().String("scan-id", "", "The unique identifier of the scan to retrieve.")
	inspector2_getCodeSecurityScanCmd.MarkFlagRequired("resource")
	inspector2_getCodeSecurityScanCmd.MarkFlagRequired("scan-id")
	inspector2Cmd.AddCommand(inspector2_getCodeSecurityScanCmd)
}
