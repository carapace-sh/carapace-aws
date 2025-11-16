package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudformation_deactivateTypeCmd = &cobra.Command{
	Use:   "deactivate-type",
	Short: "Deactivates a public third-party extension, such as a resource or module, or a CloudFormation Hook when you no longer use it.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudformation_deactivateTypeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudformation_deactivateTypeCmd).Standalone()

		cloudformation_deactivateTypeCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) for the extension in this account and Region.")
		cloudformation_deactivateTypeCmd.Flags().String("type", "", "The extension type.")
		cloudformation_deactivateTypeCmd.Flags().String("type-name", "", "The type name of the extension in this account and Region.")
	})
	cloudformationCmd.AddCommand(cloudformation_deactivateTypeCmd)
}
