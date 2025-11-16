package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspector2_listCodeSecurityScanConfigurationAssociationsCmd = &cobra.Command{
	Use:   "list-code-security-scan-configuration-associations",
	Short: "Lists the associations between code repositories and Amazon Inspector code security scan configurations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspector2_listCodeSecurityScanConfigurationAssociationsCmd).Standalone()

	inspector2_listCodeSecurityScanConfigurationAssociationsCmd.Flags().String("max-results", "", "The maximum number of results to return in the response.")
	inspector2_listCodeSecurityScanConfigurationAssociationsCmd.Flags().String("next-token", "", "A token to use for paginating results that are returned in the response.")
	inspector2_listCodeSecurityScanConfigurationAssociationsCmd.Flags().String("scan-configuration-arn", "", "The Amazon Resource Name (ARN) of the scan configuration to list associations for.")
	inspector2_listCodeSecurityScanConfigurationAssociationsCmd.MarkFlagRequired("scan-configuration-arn")
	inspector2Cmd.AddCommand(inspector2_listCodeSecurityScanConfigurationAssociationsCmd)
}
