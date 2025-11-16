package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexModels_createSlotTypeVersionCmd = &cobra.Command{
	Use:   "create-slot-type-version",
	Short: "Creates a new version of a slot type based on the `$LATEST` version of the specified slot type.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexModels_createSlotTypeVersionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lexModels_createSlotTypeVersionCmd).Standalone()

		lexModels_createSlotTypeVersionCmd.Flags().String("checksum", "", "Checksum for the `$LATEST` version of the slot type that you want to publish.")
		lexModels_createSlotTypeVersionCmd.Flags().String("name", "", "The name of the slot type that you want to create a new version for.")
		lexModels_createSlotTypeVersionCmd.MarkFlagRequired("name")
	})
	lexModelsCmd.AddCommand(lexModels_createSlotTypeVersionCmd)
}
