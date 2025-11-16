package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var health_describeAffectedEntitiesCmd = &cobra.Command{
	Use:   "describe-affected-entities",
	Short: "Returns a list of entities that have been affected by the specified events, based on the specified filter criteria.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(health_describeAffectedEntitiesCmd).Standalone()

	health_describeAffectedEntitiesCmd.Flags().String("filter", "", "Values to narrow the results returned.")
	health_describeAffectedEntitiesCmd.Flags().String("locale", "", "The locale (language) to return information in.")
	health_describeAffectedEntitiesCmd.Flags().String("max-results", "", "The maximum number of items to return in one batch, between 10 and 100, inclusive.")
	health_describeAffectedEntitiesCmd.Flags().String("next-token", "", "If the results of a search are large, only a portion of the results are returned, and a `nextToken` pagination token is returned in the response.")
	health_describeAffectedEntitiesCmd.MarkFlagRequired("filter")
	healthCmd.AddCommand(health_describeAffectedEntitiesCmd)
}
