package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexv2Models_listAggregatedUtterancesCmd = &cobra.Command{
	Use:   "list-aggregated-utterances",
	Short: "Provides a list of utterances that users have sent to the bot.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexv2Models_listAggregatedUtterancesCmd).Standalone()

	lexv2Models_listAggregatedUtterancesCmd.Flags().String("aggregation-duration", "", "The time window for aggregating the utterance information.")
	lexv2Models_listAggregatedUtterancesCmd.Flags().String("bot-alias-id", "", "The identifier of the bot alias associated with this request.")
	lexv2Models_listAggregatedUtterancesCmd.Flags().String("bot-id", "", "The unique identifier of the bot associated with this request.")
	lexv2Models_listAggregatedUtterancesCmd.Flags().String("bot-version", "", "The identifier of the bot version associated with this request.")
	lexv2Models_listAggregatedUtterancesCmd.Flags().String("filters", "", "Provides the specification of a filter used to limit the utterances in the response to only those that match the filter specification.")
	lexv2Models_listAggregatedUtterancesCmd.Flags().String("locale-id", "", "The identifier of the language and locale where the utterances were collected.")
	lexv2Models_listAggregatedUtterancesCmd.Flags().String("max-results", "", "The maximum number of utterances to return in each page of results.")
	lexv2Models_listAggregatedUtterancesCmd.Flags().String("next-token", "", "If the response from the `ListAggregatedUtterances` operation contains more results that specified in the `maxResults` parameter, a token is returned in the response.")
	lexv2Models_listAggregatedUtterancesCmd.Flags().String("sort-by", "", "Specifies sorting parameters for the list of utterances.")
	lexv2Models_listAggregatedUtterancesCmd.MarkFlagRequired("aggregation-duration")
	lexv2Models_listAggregatedUtterancesCmd.MarkFlagRequired("bot-id")
	lexv2Models_listAggregatedUtterancesCmd.MarkFlagRequired("locale-id")
	lexv2ModelsCmd.AddCommand(lexv2Models_listAggregatedUtterancesCmd)
}
