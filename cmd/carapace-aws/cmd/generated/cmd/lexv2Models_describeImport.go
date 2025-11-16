package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexv2Models_describeImportCmd = &cobra.Command{
	Use:   "describe-import",
	Short: "Gets information about a specific import.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexv2Models_describeImportCmd).Standalone()

	lexv2Models_describeImportCmd.Flags().String("import-id", "", "The unique identifier of the import to describe.")
	lexv2Models_describeImportCmd.MarkFlagRequired("import-id")
	lexv2ModelsCmd.AddCommand(lexv2Models_describeImportCmd)
}
