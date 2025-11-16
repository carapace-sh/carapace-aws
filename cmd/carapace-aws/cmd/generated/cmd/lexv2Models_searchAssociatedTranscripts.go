package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexv2Models_searchAssociatedTranscriptsCmd = &cobra.Command{
	Use:   "search-associated-transcripts",
	Short: "Search for associated transcripts that meet the specified criteria.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexv2Models_searchAssociatedTranscriptsCmd).Standalone()

	lexv2Models_searchAssociatedTranscriptsCmd.Flags().String("bot-id", "", "The unique identifier of the bot associated with the transcripts that you are searching.")
	lexv2Models_searchAssociatedTranscriptsCmd.Flags().String("bot-recommendation-id", "", "The unique identifier of the bot recommendation associated with the transcripts to search.")
	lexv2Models_searchAssociatedTranscriptsCmd.Flags().String("bot-version", "", "The version of the bot containing the transcripts that you are searching.")
	lexv2Models_searchAssociatedTranscriptsCmd.Flags().String("filters", "", "A list of filter objects.")
	lexv2Models_searchAssociatedTranscriptsCmd.Flags().String("locale-id", "", "The identifier of the language and locale of the transcripts to search.")
	lexv2Models_searchAssociatedTranscriptsCmd.Flags().String("max-results", "", "The maximum number of bot recommendations to return in each page of results.")
	lexv2Models_searchAssociatedTranscriptsCmd.Flags().String("next-index", "", "If the response from the SearchAssociatedTranscriptsRequest operation contains more results than specified in the maxResults parameter, an index is returned in the response.")
	lexv2Models_searchAssociatedTranscriptsCmd.Flags().String("search-order", "", "How SearchResults are ordered.")
	lexv2Models_searchAssociatedTranscriptsCmd.MarkFlagRequired("bot-id")
	lexv2Models_searchAssociatedTranscriptsCmd.MarkFlagRequired("bot-recommendation-id")
	lexv2Models_searchAssociatedTranscriptsCmd.MarkFlagRequired("bot-version")
	lexv2Models_searchAssociatedTranscriptsCmd.MarkFlagRequired("filters")
	lexv2Models_searchAssociatedTranscriptsCmd.MarkFlagRequired("locale-id")
	lexv2ModelsCmd.AddCommand(lexv2Models_searchAssociatedTranscriptsCmd)
}
