package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspector2_updateCodeSecurityIntegrationCmd = &cobra.Command{
	Use:   "update-code-security-integration",
	Short: "Updates an existing code security integration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspector2_updateCodeSecurityIntegrationCmd).Standalone()

	inspector2_updateCodeSecurityIntegrationCmd.Flags().String("details", "", "The updated integration details specific to the repository provider type.")
	inspector2_updateCodeSecurityIntegrationCmd.Flags().String("integration-arn", "", "The Amazon Resource Name (ARN) of the code security integration to update.")
	inspector2_updateCodeSecurityIntegrationCmd.MarkFlagRequired("details")
	inspector2_updateCodeSecurityIntegrationCmd.MarkFlagRequired("integration-arn")
	inspector2Cmd.AddCommand(inspector2_updateCodeSecurityIntegrationCmd)
}
