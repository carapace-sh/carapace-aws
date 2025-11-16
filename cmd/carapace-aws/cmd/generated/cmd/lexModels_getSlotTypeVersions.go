package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexModels_getSlotTypeVersionsCmd = &cobra.Command{
	Use:   "get-slot-type-versions",
	Short: "Gets information about all versions of a slot type.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexModels_getSlotTypeVersionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lexModels_getSlotTypeVersionsCmd).Standalone()

		lexModels_getSlotTypeVersionsCmd.Flags().String("max-results", "", "The maximum number of slot type versions to return in the response.")
		lexModels_getSlotTypeVersionsCmd.Flags().String("name", "", "The name of the slot type for which versions should be returned.")
		lexModels_getSlotTypeVersionsCmd.Flags().String("next-token", "", "A pagination token for fetching the next page of slot type versions.")
		lexModels_getSlotTypeVersionsCmd.MarkFlagRequired("name")
	})
	lexModelsCmd.AddCommand(lexModels_getSlotTypeVersionsCmd)
}
