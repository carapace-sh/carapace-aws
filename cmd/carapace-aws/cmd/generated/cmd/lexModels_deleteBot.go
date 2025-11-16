package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexModels_deleteBotCmd = &cobra.Command{
	Use:   "delete-bot",
	Short: "Deletes all versions of the bot, including the `$LATEST` version.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexModels_deleteBotCmd).Standalone()

	lexModels_deleteBotCmd.Flags().String("name", "", "The name of the bot.")
	lexModels_deleteBotCmd.MarkFlagRequired("name")
	lexModelsCmd.AddCommand(lexModels_deleteBotCmd)
}
