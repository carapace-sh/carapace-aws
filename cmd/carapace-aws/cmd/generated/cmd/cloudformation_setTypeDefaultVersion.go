package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudformation_setTypeDefaultVersionCmd = &cobra.Command{
	Use:   "set-type-default-version",
	Short: "Specify the default version of an extension.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudformation_setTypeDefaultVersionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudformation_setTypeDefaultVersionCmd).Standalone()

		cloudformation_setTypeDefaultVersionCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) of the extension for which you want version summary information.")
		cloudformation_setTypeDefaultVersionCmd.Flags().String("type", "", "The kind of extension.")
		cloudformation_setTypeDefaultVersionCmd.Flags().String("type-name", "", "The name of the extension.")
		cloudformation_setTypeDefaultVersionCmd.Flags().String("version-id", "", "The ID of a specific version of the extension.")
	})
	cloudformationCmd.AddCommand(cloudformation_setTypeDefaultVersionCmd)
}
