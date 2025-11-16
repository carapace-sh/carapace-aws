package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexModels_deleteSlotTypeVersionCmd = &cobra.Command{
	Use:   "delete-slot-type-version",
	Short: "Deletes a specific version of a slot type.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexModels_deleteSlotTypeVersionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lexModels_deleteSlotTypeVersionCmd).Standalone()

		lexModels_deleteSlotTypeVersionCmd.Flags().String("name", "", "The name of the slot type.")
		lexModels_deleteSlotTypeVersionCmd.Flags().String("version", "", "The version of the slot type to delete.")
		lexModels_deleteSlotTypeVersionCmd.MarkFlagRequired("name")
		lexModels_deleteSlotTypeVersionCmd.MarkFlagRequired("version")
	})
	lexModelsCmd.AddCommand(lexModels_deleteSlotTypeVersionCmd)
}
