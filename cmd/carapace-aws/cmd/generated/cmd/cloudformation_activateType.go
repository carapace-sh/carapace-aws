package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudformation_activateTypeCmd = &cobra.Command{
	Use:   "activate-type",
	Short: "Activates a public third-party extension, such as a resource or module, to make it available for use in stack templates in your current account and Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudformation_activateTypeCmd).Standalone()

	cloudformation_activateTypeCmd.Flags().String("auto-update", "", "Whether to automatically update the extension in this account and Region when a new *minor* version is published by the extension publisher.")
	cloudformation_activateTypeCmd.Flags().String("execution-role-arn", "", "The name of the IAM execution role to use to activate the extension.")
	cloudformation_activateTypeCmd.Flags().String("logging-config", "", "Contains logging configuration information for an extension.")
	cloudformation_activateTypeCmd.Flags().String("major-version", "", "The major version of this extension you want to activate, if multiple major versions are available.")
	cloudformation_activateTypeCmd.Flags().String("public-type-arn", "", "The Amazon Resource Name (ARN) of the public extension.")
	cloudformation_activateTypeCmd.Flags().String("publisher-id", "", "The ID of the extension publisher.")
	cloudformation_activateTypeCmd.Flags().String("type", "", "The extension type.")
	cloudformation_activateTypeCmd.Flags().String("type-name", "", "The name of the extension.")
	cloudformation_activateTypeCmd.Flags().String("type-name-alias", "", "An alias to assign to the public extension in this account and Region.")
	cloudformation_activateTypeCmd.Flags().String("version-bump", "", "Manually updates a previously-activated type to a new major or minor version, if available.")
	cloudformationCmd.AddCommand(cloudformation_activateTypeCmd)
}
