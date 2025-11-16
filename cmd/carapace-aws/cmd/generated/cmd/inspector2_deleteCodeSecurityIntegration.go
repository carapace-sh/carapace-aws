package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspector2_deleteCodeSecurityIntegrationCmd = &cobra.Command{
	Use:   "delete-code-security-integration",
	Short: "Deletes a code security integration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspector2_deleteCodeSecurityIntegrationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(inspector2_deleteCodeSecurityIntegrationCmd).Standalone()

		inspector2_deleteCodeSecurityIntegrationCmd.Flags().String("integration-arn", "", "The Amazon Resource Name (ARN) of the code security integration to delete.")
		inspector2_deleteCodeSecurityIntegrationCmd.MarkFlagRequired("integration-arn")
	})
	inspector2Cmd.AddCommand(inspector2_deleteCodeSecurityIntegrationCmd)
}
