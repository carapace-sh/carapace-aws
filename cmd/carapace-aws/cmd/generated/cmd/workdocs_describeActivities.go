package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workdocs_describeActivitiesCmd = &cobra.Command{
	Use:   "describe-activities",
	Short: "Describes the user activities in a specified time period.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workdocs_describeActivitiesCmd).Standalone()

	workdocs_describeActivitiesCmd.Flags().String("activity-types", "", "Specifies which activity types to include in the response.")
	workdocs_describeActivitiesCmd.Flags().String("authentication-token", "", "Amazon WorkDocs authentication token.")
	workdocs_describeActivitiesCmd.Flags().String("end-time", "", "The timestamp that determines the end time of the activities.")
	workdocs_describeActivitiesCmd.Flags().String("include-indirect-activities", "", "Includes indirect activities.")
	workdocs_describeActivitiesCmd.Flags().String("limit", "", "The maximum number of items to return.")
	workdocs_describeActivitiesCmd.Flags().String("marker", "", "The marker for the next set of results.")
	workdocs_describeActivitiesCmd.Flags().String("organization-id", "", "The ID of the organization.")
	workdocs_describeActivitiesCmd.Flags().String("resource-id", "", "The document or folder ID for which to describe activity types.")
	workdocs_describeActivitiesCmd.Flags().String("start-time", "", "The timestamp that determines the starting time of the activities.")
	workdocs_describeActivitiesCmd.Flags().String("user-id", "", "The ID of the user who performed the action.")
	workdocsCmd.AddCommand(workdocs_describeActivitiesCmd)
}
