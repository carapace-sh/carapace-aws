package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var health_describeEntityAggregatesCmd = &cobra.Command{
	Use:   "describe-entity-aggregates",
	Short: "Returns the number of entities that are affected by each of the specified events.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(health_describeEntityAggregatesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(health_describeEntityAggregatesCmd).Standalone()

		health_describeEntityAggregatesCmd.Flags().String("event-arns", "", "A list of event ARNs (unique identifiers).")
	})
	healthCmd.AddCommand(health_describeEntityAggregatesCmd)
}
