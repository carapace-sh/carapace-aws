package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_describeScheduledActionsCmd = &cobra.Command{
	Use:   "describe-scheduled-actions",
	Short: "Describes properties of scheduled actions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_describeScheduledActionsCmd).Standalone()

	redshift_describeScheduledActionsCmd.Flags().String("active", "", "If true, retrieve only active scheduled actions.")
	redshift_describeScheduledActionsCmd.Flags().String("end-time", "", "The end time in UTC of the scheduled action to retrieve.")
	redshift_describeScheduledActionsCmd.Flags().String("filters", "", "List of scheduled action filters.")
	redshift_describeScheduledActionsCmd.Flags().String("marker", "", "An optional parameter that specifies the starting point to return a set of response records.")
	redshift_describeScheduledActionsCmd.Flags().String("max-records", "", "The maximum number of response records to return in each call.")
	redshift_describeScheduledActionsCmd.Flags().String("scheduled-action-name", "", "The name of the scheduled action to retrieve.")
	redshift_describeScheduledActionsCmd.Flags().String("start-time", "", "The start time in UTC of the scheduled actions to retrieve.")
	redshift_describeScheduledActionsCmd.Flags().String("target-action-type", "", "The type of the scheduled actions to retrieve.")
	redshiftCmd.AddCommand(redshift_describeScheduledActionsCmd)
}
