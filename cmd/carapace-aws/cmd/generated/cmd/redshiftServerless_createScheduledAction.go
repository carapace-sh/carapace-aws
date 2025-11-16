package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshiftServerless_createScheduledActionCmd = &cobra.Command{
	Use:   "create-scheduled-action",
	Short: "Creates a scheduled action.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshiftServerless_createScheduledActionCmd).Standalone()

	redshiftServerless_createScheduledActionCmd.Flags().Bool("enabled", false, "Indicates whether the schedule is enabled.")
	redshiftServerless_createScheduledActionCmd.Flags().String("end-time", "", "The end time in UTC when the schedule is no longer active.")
	redshiftServerless_createScheduledActionCmd.Flags().String("namespace-name", "", "The name of the namespace for which to create a scheduled action.")
	redshiftServerless_createScheduledActionCmd.Flags().Bool("no-enabled", false, "Indicates whether the schedule is enabled.")
	redshiftServerless_createScheduledActionCmd.Flags().String("role-arn", "", "The ARN of the IAM role to assume to run the scheduled action.")
	redshiftServerless_createScheduledActionCmd.Flags().String("schedule", "", "The schedule for a one-time (at timestamp format) or recurring (cron format) scheduled action.")
	redshiftServerless_createScheduledActionCmd.Flags().String("scheduled-action-description", "", "The description of the scheduled action.")
	redshiftServerless_createScheduledActionCmd.Flags().String("scheduled-action-name", "", "The name of the scheduled action.")
	redshiftServerless_createScheduledActionCmd.Flags().String("start-time", "", "The start time in UTC when the schedule is active.")
	redshiftServerless_createScheduledActionCmd.Flags().String("target-action", "", "")
	redshiftServerless_createScheduledActionCmd.MarkFlagRequired("namespace-name")
	redshiftServerless_createScheduledActionCmd.Flag("no-enabled").Hidden = true
	redshiftServerless_createScheduledActionCmd.MarkFlagRequired("role-arn")
	redshiftServerless_createScheduledActionCmd.MarkFlagRequired("schedule")
	redshiftServerless_createScheduledActionCmd.MarkFlagRequired("scheduled-action-name")
	redshiftServerless_createScheduledActionCmd.MarkFlagRequired("target-action")
	redshiftServerlessCmd.AddCommand(redshiftServerless_createScheduledActionCmd)
}
