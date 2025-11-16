package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_describeEventCategoriesCmd = &cobra.Command{
	Use:   "describe-event-categories",
	Short: "Displays a list of event categories for all event source types, or for a specified source type.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_describeEventCategoriesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(redshift_describeEventCategoriesCmd).Standalone()

		redshift_describeEventCategoriesCmd.Flags().String("source-type", "", "The source type, such as cluster or parameter group, to which the described event categories apply.")
	})
	redshiftCmd.AddCommand(redshift_describeEventCategoriesCmd)
}
