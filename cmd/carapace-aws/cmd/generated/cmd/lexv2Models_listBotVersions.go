package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexv2Models_listBotVersionsCmd = &cobra.Command{
	Use:   "list-bot-versions",
	Short: "Gets information about all of the versions of a bot.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexv2Models_listBotVersionsCmd).Standalone()

	lexv2Models_listBotVersionsCmd.Flags().String("bot-id", "", "The identifier of the bot to list versions for.")
	lexv2Models_listBotVersionsCmd.Flags().String("max-results", "", "The maximum number of versions to return in each page of results.")
	lexv2Models_listBotVersionsCmd.Flags().String("next-token", "", "If the response to the `ListBotVersion` operation contains more results than specified in the `maxResults` parameter, a token is returned in the response.")
	lexv2Models_listBotVersionsCmd.Flags().String("sort-by", "", "Specifies sorting parameters for the list of versions.")
	lexv2Models_listBotVersionsCmd.MarkFlagRequired("bot-id")
	lexv2ModelsCmd.AddCommand(lexv2Models_listBotVersionsCmd)
}
