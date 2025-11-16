package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspector2_createCodeSecurityIntegrationCmd = &cobra.Command{
	Use:   "create-code-security-integration",
	Short: "Creates a code security integration with a source code repository provider.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspector2_createCodeSecurityIntegrationCmd).Standalone()

	inspector2_createCodeSecurityIntegrationCmd.Flags().String("details", "", "The integration details specific to the repository provider type.")
	inspector2_createCodeSecurityIntegrationCmd.Flags().String("name", "", "The name of the code security integration.")
	inspector2_createCodeSecurityIntegrationCmd.Flags().String("tags", "", "The tags to apply to the code security integration.")
	inspector2_createCodeSecurityIntegrationCmd.Flags().String("type", "", "The type of repository provider for the integration.")
	inspector2_createCodeSecurityIntegrationCmd.MarkFlagRequired("name")
	inspector2_createCodeSecurityIntegrationCmd.MarkFlagRequired("type")
	inspector2Cmd.AddCommand(inspector2_createCodeSecurityIntegrationCmd)
}
