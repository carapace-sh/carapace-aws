package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexv2Models_deleteImportCmd = &cobra.Command{
	Use:   "delete-import",
	Short: "Removes a previous import and the associated file stored in an S3 bucket.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexv2Models_deleteImportCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lexv2Models_deleteImportCmd).Standalone()

		lexv2Models_deleteImportCmd.Flags().String("import-id", "", "The unique identifier of the import to delete.")
		lexv2Models_deleteImportCmd.MarkFlagRequired("import-id")
	})
	lexv2ModelsCmd.AddCommand(lexv2Models_deleteImportCmd)
}
