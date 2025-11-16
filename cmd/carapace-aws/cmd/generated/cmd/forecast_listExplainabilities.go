package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var forecast_listExplainabilitiesCmd = &cobra.Command{
	Use:   "list-explainabilities",
	Short: "Returns a list of Explainability resources created using the [CreateExplainability]() operation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(forecast_listExplainabilitiesCmd).Standalone()

	forecast_listExplainabilitiesCmd.Flags().String("filters", "", "An array of filters.")
	forecast_listExplainabilitiesCmd.Flags().String("max-results", "", "The number of items returned in the response.")
	forecast_listExplainabilitiesCmd.Flags().String("next-token", "", "If the result of the previous request was truncated, the response includes a NextToken.")
	forecastCmd.AddCommand(forecast_listExplainabilitiesCmd)
}
