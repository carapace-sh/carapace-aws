package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexModels_deleteBotVersionCmd = &cobra.Command{
	Use:   "delete-bot-version",
	Short: "Deletes a specific version of a bot.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexModels_deleteBotVersionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lexModels_deleteBotVersionCmd).Standalone()

		lexModels_deleteBotVersionCmd.Flags().String("name", "", "The name of the bot.")
		lexModels_deleteBotVersionCmd.Flags().String("version", "", "The version of the bot to delete.")
		lexModels_deleteBotVersionCmd.MarkFlagRequired("name")
		lexModels_deleteBotVersionCmd.MarkFlagRequired("version")
	})
	lexModelsCmd.AddCommand(lexModels_deleteBotVersionCmd)
}
