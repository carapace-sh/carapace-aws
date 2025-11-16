package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var freetier_listAccountActivitiesCmd = &cobra.Command{
	Use:   "list-account-activities",
	Short: "Returns a list of activities that are available.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(freetier_listAccountActivitiesCmd).Standalone()

	freetier_listAccountActivitiesCmd.Flags().String("filter-activity-statuses", "", "The activity status filter.")
	freetier_listAccountActivitiesCmd.Flags().String("language-code", "", "The language code used to return translated titles.")
	freetier_listAccountActivitiesCmd.Flags().String("max-results", "", "The maximum number of items to return for this request.")
	freetier_listAccountActivitiesCmd.Flags().String("next-token", "", "A token from a previous paginated response.")
	freetierCmd.AddCommand(freetier_listAccountActivitiesCmd)
}
