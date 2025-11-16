package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptune_describeEventCategoriesCmd = &cobra.Command{
	Use:   "describe-event-categories",
	Short: "Displays a list of categories for all event source types, or, if specified, for a specified source type.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptune_describeEventCategoriesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(neptune_describeEventCategoriesCmd).Standalone()

		neptune_describeEventCategoriesCmd.Flags().String("filters", "", "This parameter is not currently supported.")
		neptune_describeEventCategoriesCmd.Flags().String("source-type", "", "The type of source that is generating the events.")
	})
	neptuneCmd.AddCommand(neptune_describeEventCategoriesCmd)
}
