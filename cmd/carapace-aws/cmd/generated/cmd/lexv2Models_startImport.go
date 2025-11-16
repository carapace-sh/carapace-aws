package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexv2Models_startImportCmd = &cobra.Command{
	Use:   "start-import",
	Short: "Starts importing a bot, bot locale, or custom vocabulary from a zip archive that you uploaded to an S3 bucket.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexv2Models_startImportCmd).Standalone()

	lexv2Models_startImportCmd.Flags().String("file-password", "", "The password used to encrypt the zip archive that contains the resource definition.")
	lexv2Models_startImportCmd.Flags().String("import-id", "", "The unique identifier for the import.")
	lexv2Models_startImportCmd.Flags().String("merge-strategy", "", "The strategy to use when there is a name conflict between the imported resource and an existing resource.")
	lexv2Models_startImportCmd.Flags().String("resource-specification", "", "Parameters for creating the bot, bot locale or custom vocabulary.")
	lexv2Models_startImportCmd.MarkFlagRequired("import-id")
	lexv2Models_startImportCmd.MarkFlagRequired("merge-strategy")
	lexv2Models_startImportCmd.MarkFlagRequired("resource-specification")
	lexv2ModelsCmd.AddCommand(lexv2Models_startImportCmd)
}
