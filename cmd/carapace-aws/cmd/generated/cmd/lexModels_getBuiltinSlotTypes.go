package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexModels_getBuiltinSlotTypesCmd = &cobra.Command{
	Use:   "get-builtin-slot-types",
	Short: "Gets a list of built-in slot types that meet the specified criteria.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexModels_getBuiltinSlotTypesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lexModels_getBuiltinSlotTypesCmd).Standalone()

		lexModels_getBuiltinSlotTypesCmd.Flags().String("locale", "", "A list of locales that the slot type supports.")
		lexModels_getBuiltinSlotTypesCmd.Flags().String("max-results", "", "The maximum number of slot types to return in the response.")
		lexModels_getBuiltinSlotTypesCmd.Flags().String("next-token", "", "A pagination token that fetches the next page of slot types.")
		lexModels_getBuiltinSlotTypesCmd.Flags().String("signature-contains", "", "Substring to match in built-in slot type signatures.")
	})
	lexModelsCmd.AddCommand(lexModels_getBuiltinSlotTypesCmd)
}
