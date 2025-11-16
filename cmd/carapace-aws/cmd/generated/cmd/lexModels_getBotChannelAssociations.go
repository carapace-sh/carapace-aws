package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexModels_getBotChannelAssociationsCmd = &cobra.Command{
	Use:   "get-bot-channel-associations",
	Short: "Returns a list of all of the channels associated with the specified bot.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexModels_getBotChannelAssociationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lexModels_getBotChannelAssociationsCmd).Standalone()

		lexModels_getBotChannelAssociationsCmd.Flags().String("bot-alias", "", "An alias pointing to the specific version of the Amazon Lex bot to which this association is being made.")
		lexModels_getBotChannelAssociationsCmd.Flags().String("bot-name", "", "The name of the Amazon Lex bot in the association.")
		lexModels_getBotChannelAssociationsCmd.Flags().String("max-results", "", "The maximum number of associations to return in the response.")
		lexModels_getBotChannelAssociationsCmd.Flags().String("name-contains", "", "Substring to match in channel association names.")
		lexModels_getBotChannelAssociationsCmd.Flags().String("next-token", "", "A pagination token for fetching the next page of associations.")
		lexModels_getBotChannelAssociationsCmd.MarkFlagRequired("bot-alias")
		lexModels_getBotChannelAssociationsCmd.MarkFlagRequired("bot-name")
	})
	lexModelsCmd.AddCommand(lexModels_getBotChannelAssociationsCmd)
}
