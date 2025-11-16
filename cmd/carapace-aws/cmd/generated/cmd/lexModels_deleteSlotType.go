package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexModels_deleteSlotTypeCmd = &cobra.Command{
	Use:   "delete-slot-type",
	Short: "Deletes all versions of the slot type, including the `$LATEST` version.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexModels_deleteSlotTypeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lexModels_deleteSlotTypeCmd).Standalone()

		lexModels_deleteSlotTypeCmd.Flags().String("name", "", "The name of the slot type.")
		lexModels_deleteSlotTypeCmd.MarkFlagRequired("name")
	})
	lexModelsCmd.AddCommand(lexModels_deleteSlotTypeCmd)
}
