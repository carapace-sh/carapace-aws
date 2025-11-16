package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dms_describeEventCategoriesCmd = &cobra.Command{
	Use:   "describe-event-categories",
	Short: "Lists categories for all event source types, or, if specified, for a specified source type.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dms_describeEventCategoriesCmd).Standalone()

	dms_describeEventCategoriesCmd.Flags().String("filters", "", "Filters applied to the event categories.")
	dms_describeEventCategoriesCmd.Flags().String("source-type", "", "The type of DMS resource that generates events.")
	dmsCmd.AddCommand(dms_describeEventCategoriesCmd)
}
