package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspector2_listCodeSecurityScanConfigurationsCmd = &cobra.Command{
	Use:   "list-code-security-scan-configurations",
	Short: "Lists all code security scan configurations in your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspector2_listCodeSecurityScanConfigurationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(inspector2_listCodeSecurityScanConfigurationsCmd).Standalone()

		inspector2_listCodeSecurityScanConfigurationsCmd.Flags().String("max-results", "", "The maximum number of results to return in a single call.")
		inspector2_listCodeSecurityScanConfigurationsCmd.Flags().String("next-token", "", "A token to use for paginating results that are returned in the response.")
	})
	inspector2Cmd.AddCommand(inspector2_listCodeSecurityScanConfigurationsCmd)
}
