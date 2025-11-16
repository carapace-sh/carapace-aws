package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var health_describeEventsCmd = &cobra.Command{
	Use:   "describe-events",
	Short: "Returns information about events that meet the specified filter criteria.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(health_describeEventsCmd).Standalone()

	health_describeEventsCmd.Flags().String("filter", "", "Values to narrow the results returned.")
	health_describeEventsCmd.Flags().String("locale", "", "The locale (language) to return information in.")
	health_describeEventsCmd.Flags().String("max-results", "", "The maximum number of items to return in one batch, between 10 and 100, inclusive.")
	health_describeEventsCmd.Flags().String("next-token", "", "If the results of a search are large, only a portion of the results are returned, and a `nextToken` pagination token is returned in the response.")
	healthCmd.AddCommand(health_describeEventsCmd)
}
