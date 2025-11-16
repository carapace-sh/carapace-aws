package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_createScheduledActionCmd = &cobra.Command{
	Use:   "create-scheduled-action",
	Short: "Creates a scheduled action.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_createScheduledActionCmd).Standalone()

	redshift_createScheduledActionCmd.Flags().String("enable", "", "If true, the schedule is enabled.")
	redshift_createScheduledActionCmd.Flags().String("end-time", "", "The end time in UTC of the scheduled action.")
	redshift_createScheduledActionCmd.Flags().String("iam-role", "", "The IAM role to assume to run the target action.")
	redshift_createScheduledActionCmd.Flags().String("schedule", "", "The schedule in `at( )` or `cron( )` format.")
	redshift_createScheduledActionCmd.Flags().String("scheduled-action-description", "", "The description of the scheduled action.")
	redshift_createScheduledActionCmd.Flags().String("scheduled-action-name", "", "The name of the scheduled action.")
	redshift_createScheduledActionCmd.Flags().String("start-time", "", "The start time in UTC of the scheduled action.")
	redshift_createScheduledActionCmd.Flags().String("target-action", "", "A JSON format string of the Amazon Redshift API operation with input parameters.")
	redshift_createScheduledActionCmd.MarkFlagRequired("iam-role")
	redshift_createScheduledActionCmd.MarkFlagRequired("schedule")
	redshift_createScheduledActionCmd.MarkFlagRequired("scheduled-action-name")
	redshift_createScheduledActionCmd.MarkFlagRequired("target-action")
	redshiftCmd.AddCommand(redshift_createScheduledActionCmd)
}
