package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexModels_getBotChannelAssociationCmd = &cobra.Command{
	Use:   "get-bot-channel-association",
	Short: "Returns information about the association between an Amazon Lex bot and a messaging platform.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexModels_getBotChannelAssociationCmd).Standalone()

	lexModels_getBotChannelAssociationCmd.Flags().String("bot-alias", "", "An alias pointing to the specific version of the Amazon Lex bot to which this association is being made.")
	lexModels_getBotChannelAssociationCmd.Flags().String("bot-name", "", "The name of the Amazon Lex bot.")
	lexModels_getBotChannelAssociationCmd.Flags().String("name", "", "The name of the association between the bot and the channel.")
	lexModels_getBotChannelAssociationCmd.MarkFlagRequired("bot-alias")
	lexModels_getBotChannelAssociationCmd.MarkFlagRequired("bot-name")
	lexModels_getBotChannelAssociationCmd.MarkFlagRequired("name")
	lexModelsCmd.AddCommand(lexModels_getBotChannelAssociationCmd)
}
