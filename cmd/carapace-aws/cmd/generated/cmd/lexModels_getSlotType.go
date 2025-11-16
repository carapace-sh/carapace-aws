package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexModels_getSlotTypeCmd = &cobra.Command{
	Use:   "get-slot-type",
	Short: "Returns information about a specific version of a slot type.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexModels_getSlotTypeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lexModels_getSlotTypeCmd).Standalone()

		lexModels_getSlotTypeCmd.Flags().String("name", "", "The name of the slot type.")
		lexModels_getSlotTypeCmd.Flags().String("version", "", "The version of the slot type.")
		lexModels_getSlotTypeCmd.MarkFlagRequired("name")
		lexModels_getSlotTypeCmd.MarkFlagRequired("version")
	})
	lexModelsCmd.AddCommand(lexModels_getSlotTypeCmd)
}
