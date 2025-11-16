package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var guardduty_createFilterCmd = &cobra.Command{
	Use:   "create-filter",
	Short: "Creates a filter using the specified finding criteria.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(guardduty_createFilterCmd).Standalone()

	guardduty_createFilterCmd.Flags().String("action", "", "Specifies the action that is to be applied to the findings that match the filter.")
	guardduty_createFilterCmd.Flags().String("client-token", "", "The idempotency token for the create request.")
	guardduty_createFilterCmd.Flags().String("description", "", "The description of the filter.")
	guardduty_createFilterCmd.Flags().String("detector-id", "", "The detector ID associated with the GuardDuty account for which you want to create a filter.")
	guardduty_createFilterCmd.Flags().String("finding-criteria", "", "Represents the criteria to be used in the filter for querying findings.")
	guardduty_createFilterCmd.Flags().String("name", "", "The name of the filter.")
	guardduty_createFilterCmd.Flags().String("rank", "", "Specifies the position of the filter in the list of current filters.")
	guardduty_createFilterCmd.Flags().String("tags", "", "The tags to be added to a new filter resource.")
	guardduty_createFilterCmd.MarkFlagRequired("detector-id")
	guardduty_createFilterCmd.MarkFlagRequired("finding-criteria")
	guardduty_createFilterCmd.MarkFlagRequired("name")
	guarddutyCmd.AddCommand(guardduty_createFilterCmd)
}
