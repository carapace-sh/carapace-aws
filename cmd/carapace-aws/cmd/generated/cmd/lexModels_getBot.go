package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexModels_getBotCmd = &cobra.Command{
	Use:   "get-bot",
	Short: "Returns metadata information for a specific bot.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexModels_getBotCmd).Standalone()

	lexModels_getBotCmd.Flags().String("name", "", "The name of the bot.")
	lexModels_getBotCmd.Flags().String("version-or-alias", "", "The version or alias of the bot.")
	lexModels_getBotCmd.MarkFlagRequired("name")
	lexModels_getBotCmd.MarkFlagRequired("version-or-alias")
	lexModelsCmd.AddCommand(lexModels_getBotCmd)
}
