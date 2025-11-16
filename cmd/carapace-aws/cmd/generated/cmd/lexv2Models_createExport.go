package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexv2Models_createExportCmd = &cobra.Command{
	Use:   "create-export",
	Short: "Creates a zip archive containing the contents of a bot or a bot locale.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexv2Models_createExportCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lexv2Models_createExportCmd).Standalone()

		lexv2Models_createExportCmd.Flags().String("file-format", "", "The file format of the bot or bot locale definition files.")
		lexv2Models_createExportCmd.Flags().String("file-password", "", "An password to use to encrypt the exported archive.")
		lexv2Models_createExportCmd.Flags().String("resource-specification", "", "Specifies the type of resource to export, either a bot or a bot locale.")
		lexv2Models_createExportCmd.MarkFlagRequired("file-format")
		lexv2Models_createExportCmd.MarkFlagRequired("resource-specification")
	})
	lexv2ModelsCmd.AddCommand(lexv2Models_createExportCmd)
}
