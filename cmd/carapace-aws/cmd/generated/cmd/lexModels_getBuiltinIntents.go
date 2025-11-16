package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexModels_getBuiltinIntentsCmd = &cobra.Command{
	Use:   "get-builtin-intents",
	Short: "Gets a list of built-in intents that meet the specified criteria.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexModels_getBuiltinIntentsCmd).Standalone()

	lexModels_getBuiltinIntentsCmd.Flags().String("locale", "", "A list of locales that the intent supports.")
	lexModels_getBuiltinIntentsCmd.Flags().String("max-results", "", "The maximum number of intents to return in the response.")
	lexModels_getBuiltinIntentsCmd.Flags().String("next-token", "", "A pagination token that fetches the next page of intents.")
	lexModels_getBuiltinIntentsCmd.Flags().String("signature-contains", "", "Substring to match in built-in intent signatures.")
	lexModelsCmd.AddCommand(lexModels_getBuiltinIntentsCmd)
}
