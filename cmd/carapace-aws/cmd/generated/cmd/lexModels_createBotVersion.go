package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexModels_createBotVersionCmd = &cobra.Command{
	Use:   "create-bot-version",
	Short: "Creates a new version of the bot based on the `$LATEST` version.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexModels_createBotVersionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lexModels_createBotVersionCmd).Standalone()

		lexModels_createBotVersionCmd.Flags().String("checksum", "", "Identifies a specific revision of the `$LATEST` version of the bot.")
		lexModels_createBotVersionCmd.Flags().String("name", "", "The name of the bot that you want to create a new version of.")
		lexModels_createBotVersionCmd.MarkFlagRequired("name")
	})
	lexModelsCmd.AddCommand(lexModels_createBotVersionCmd)
}
