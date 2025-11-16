package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexModels_getIntentsCmd = &cobra.Command{
	Use:   "get-intents",
	Short: "Returns intent information as follows:\n\n- If you specify the `nameContains` field, returns the `$LATEST` version of all intents that contain the specified string.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexModels_getIntentsCmd).Standalone()

	lexModels_getIntentsCmd.Flags().String("max-results", "", "The maximum number of intents to return in the response.")
	lexModels_getIntentsCmd.Flags().String("name-contains", "", "Substring to match in intent names.")
	lexModels_getIntentsCmd.Flags().String("next-token", "", "A pagination token that fetches the next page of intents.")
	lexModelsCmd.AddCommand(lexModels_getIntentsCmd)
}
