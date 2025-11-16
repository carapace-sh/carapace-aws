package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudformation_describeTypeCmd = &cobra.Command{
	Use:   "describe-type",
	Short: "Returns detailed information about an extension from the CloudFormation registry in your current account and Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudformation_describeTypeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudformation_describeTypeCmd).Standalone()

		cloudformation_describeTypeCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) of the extension.")
		cloudformation_describeTypeCmd.Flags().String("public-version-number", "", "The version number of a public third-party extension.")
		cloudformation_describeTypeCmd.Flags().String("publisher-id", "", "The publisher ID of the extension publisher.")
		cloudformation_describeTypeCmd.Flags().String("type", "", "The kind of extension.")
		cloudformation_describeTypeCmd.Flags().String("type-name", "", "The name of the extension.")
		cloudformation_describeTypeCmd.Flags().String("version-id", "", "The ID of a specific version of the extension.")
	})
	cloudformationCmd.AddCommand(cloudformation_describeTypeCmd)
}
