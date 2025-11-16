package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexModels_getSlotTypesCmd = &cobra.Command{
	Use:   "get-slot-types",
	Short: "Returns slot type information as follows:\n\n- If you specify the `nameContains` field, returns the `$LATEST` version of all slot types that contain the specified string.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexModels_getSlotTypesCmd).Standalone()

	lexModels_getSlotTypesCmd.Flags().String("max-results", "", "The maximum number of slot types to return in the response.")
	lexModels_getSlotTypesCmd.Flags().String("name-contains", "", "Substring to match in slot type names.")
	lexModels_getSlotTypesCmd.Flags().String("next-token", "", "A pagination token that fetches the next page of slot types.")
	lexModelsCmd.AddCommand(lexModels_getSlotTypesCmd)
}
