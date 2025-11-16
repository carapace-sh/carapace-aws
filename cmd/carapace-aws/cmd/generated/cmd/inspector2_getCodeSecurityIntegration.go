package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspector2_getCodeSecurityIntegrationCmd = &cobra.Command{
	Use:   "get-code-security-integration",
	Short: "Retrieves information about a code security integration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspector2_getCodeSecurityIntegrationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(inspector2_getCodeSecurityIntegrationCmd).Standalone()

		inspector2_getCodeSecurityIntegrationCmd.Flags().String("integration-arn", "", "The Amazon Resource Name (ARN) of the code security integration to retrieve.")
		inspector2_getCodeSecurityIntegrationCmd.Flags().String("tags", "", "The tags associated with the code security integration.")
		inspector2_getCodeSecurityIntegrationCmd.MarkFlagRequired("integration-arn")
	})
	inspector2Cmd.AddCommand(inspector2_getCodeSecurityIntegrationCmd)
}
