package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudformation_deregisterTypeCmd = &cobra.Command{
	Use:   "deregister-type",
	Short: "Marks an extension or extension version as `DEPRECATED` in the CloudFormation registry, removing it from active use.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudformation_deregisterTypeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudformation_deregisterTypeCmd).Standalone()

		cloudformation_deregisterTypeCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) of the extension.")
		cloudformation_deregisterTypeCmd.Flags().String("type", "", "The kind of extension.")
		cloudformation_deregisterTypeCmd.Flags().String("type-name", "", "The name of the extension.")
		cloudformation_deregisterTypeCmd.Flags().String("version-id", "", "The ID of a specific version of the extension.")
	})
	cloudformationCmd.AddCommand(cloudformation_deregisterTypeCmd)
}
