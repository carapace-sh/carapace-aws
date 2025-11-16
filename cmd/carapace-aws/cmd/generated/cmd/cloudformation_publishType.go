package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudformation_publishTypeCmd = &cobra.Command{
	Use:   "publish-type",
	Short: "Publishes the specified extension to the CloudFormation registry as a public extension in this Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudformation_publishTypeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudformation_publishTypeCmd).Standalone()

		cloudformation_publishTypeCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) of the extension.")
		cloudformation_publishTypeCmd.Flags().String("public-version-number", "", "The version number to assign to this version of the extension.")
		cloudformation_publishTypeCmd.Flags().String("type", "", "The type of the extension.")
		cloudformation_publishTypeCmd.Flags().String("type-name", "", "The name of the extension.")
	})
	cloudformationCmd.AddCommand(cloudformation_publishTypeCmd)
}
