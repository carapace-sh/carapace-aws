package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexModels_getBotsCmd = &cobra.Command{
	Use:   "get-bots",
	Short: "Returns bot information as follows:\n\n- If you provide the `nameContains` field, the response includes information for the `$LATEST` version of all bots whose name contains the specified string.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexModels_getBotsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lexModels_getBotsCmd).Standalone()

		lexModels_getBotsCmd.Flags().String("max-results", "", "The maximum number of bots to return in the response that the request will return.")
		lexModels_getBotsCmd.Flags().String("name-contains", "", "Substring to match in bot names.")
		lexModels_getBotsCmd.Flags().String("next-token", "", "A pagination token that fetches the next page of bots.")
	})
	lexModelsCmd.AddCommand(lexModels_getBotsCmd)
}
