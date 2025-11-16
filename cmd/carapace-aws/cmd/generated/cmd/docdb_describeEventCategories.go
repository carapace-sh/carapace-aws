package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var docdb_describeEventCategoriesCmd = &cobra.Command{
	Use:   "describe-event-categories",
	Short: "Displays a list of categories for all event source types, or, if specified, for a specified source type.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(docdb_describeEventCategoriesCmd).Standalone()

	docdb_describeEventCategoriesCmd.Flags().String("filters", "", "This parameter is not currently supported.")
	docdb_describeEventCategoriesCmd.Flags().String("source-type", "", "The type of source that is generating the events.")
	docdbCmd.AddCommand(docdb_describeEventCategoriesCmd)
}
