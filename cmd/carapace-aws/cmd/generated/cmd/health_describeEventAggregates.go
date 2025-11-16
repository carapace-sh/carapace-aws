package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var health_describeEventAggregatesCmd = &cobra.Command{
	Use:   "describe-event-aggregates",
	Short: "Returns the number of events of each event type (issue, scheduled change, and account notification).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(health_describeEventAggregatesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(health_describeEventAggregatesCmd).Standalone()

		health_describeEventAggregatesCmd.Flags().String("aggregate-field", "", "The only currently supported value is `eventTypeCategory`.")
		health_describeEventAggregatesCmd.Flags().String("filter", "", "Values to narrow the results returned.")
		health_describeEventAggregatesCmd.Flags().String("max-results", "", "The maximum number of items to return in one batch, between 10 and 100, inclusive.")
		health_describeEventAggregatesCmd.Flags().String("next-token", "", "If the results of a search are large, only a portion of the results are returned, and a `nextToken` pagination token is returned in the response.")
		health_describeEventAggregatesCmd.MarkFlagRequired("aggregate-field")
	})
	healthCmd.AddCommand(health_describeEventAggregatesCmd)
}
