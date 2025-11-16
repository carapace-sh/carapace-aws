package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var health_describeEntityAggregatesForOrganizationCmd = &cobra.Command{
	Use:   "describe-entity-aggregates-for-organization",
	Short: "Returns a list of entity aggregates for your Organizations that are affected by each of the specified events.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(health_describeEntityAggregatesForOrganizationCmd).Standalone()

	health_describeEntityAggregatesForOrganizationCmd.Flags().String("aws-account-ids", "", "A list of 12-digit Amazon Web Services account numbers that contains the affected entities.")
	health_describeEntityAggregatesForOrganizationCmd.Flags().String("event-arns", "", "A list of event ARNs (unique identifiers).")
	health_describeEntityAggregatesForOrganizationCmd.MarkFlagRequired("event-arns")
	healthCmd.AddCommand(health_describeEntityAggregatesForOrganizationCmd)
}
