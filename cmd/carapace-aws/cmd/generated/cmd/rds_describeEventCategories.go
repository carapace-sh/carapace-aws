package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_describeEventCategoriesCmd = &cobra.Command{
	Use:   "describe-event-categories",
	Short: "Displays a list of categories for all event source types, or, if specified, for a specified source type.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_describeEventCategoriesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rds_describeEventCategoriesCmd).Standalone()

		rds_describeEventCategoriesCmd.Flags().String("filters", "", "This parameter isn't currently supported.")
		rds_describeEventCategoriesCmd.Flags().String("source-type", "", "The type of source that is generating the events.")
	})
	rdsCmd.AddCommand(rds_describeEventCategoriesCmd)
}
