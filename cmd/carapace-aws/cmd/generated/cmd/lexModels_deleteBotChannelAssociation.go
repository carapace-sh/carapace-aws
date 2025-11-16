package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexModels_deleteBotChannelAssociationCmd = &cobra.Command{
	Use:   "delete-bot-channel-association",
	Short: "Deletes the association between an Amazon Lex bot and a messaging platform.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexModels_deleteBotChannelAssociationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lexModels_deleteBotChannelAssociationCmd).Standalone()

		lexModels_deleteBotChannelAssociationCmd.Flags().String("bot-alias", "", "An alias that points to the specific version of the Amazon Lex bot to which this association is being made.")
		lexModels_deleteBotChannelAssociationCmd.Flags().String("bot-name", "", "The name of the Amazon Lex bot.")
		lexModels_deleteBotChannelAssociationCmd.Flags().String("name", "", "The name of the association.")
		lexModels_deleteBotChannelAssociationCmd.MarkFlagRequired("bot-alias")
		lexModels_deleteBotChannelAssociationCmd.MarkFlagRequired("bot-name")
		lexModels_deleteBotChannelAssociationCmd.MarkFlagRequired("name")
	})
	lexModelsCmd.AddCommand(lexModels_deleteBotChannelAssociationCmd)
}
