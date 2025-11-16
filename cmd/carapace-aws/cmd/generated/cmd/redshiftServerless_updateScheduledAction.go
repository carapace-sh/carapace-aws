package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshiftServerless_updateScheduledActionCmd = &cobra.Command{
	Use:   "update-scheduled-action",
	Short: "Updates a scheduled action.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshiftServerless_updateScheduledActionCmd).Standalone()

	redshiftServerless_updateScheduledActionCmd.Flags().Bool("enabled", false, "Specifies whether to enable the scheduled action.")
	redshiftServerless_updateScheduledActionCmd.Flags().String("end-time", "", "The end time in UTC of the scheduled action to update.")
	redshiftServerless_updateScheduledActionCmd.Flags().Bool("no-enabled", false, "Specifies whether to enable the scheduled action.")
	redshiftServerless_updateScheduledActionCmd.Flags().String("role-arn", "", "The ARN of the IAM role to assume to run the scheduled action.")
	redshiftServerless_updateScheduledActionCmd.Flags().String("schedule", "", "The schedule for a one-time (at timestamp format) or recurring (cron format) scheduled action.")
	redshiftServerless_updateScheduledActionCmd.Flags().String("scheduled-action-description", "", "The descripion of the scheduled action to update to.")
	redshiftServerless_updateScheduledActionCmd.Flags().String("scheduled-action-name", "", "The name of the scheduled action to update to.")
	redshiftServerless_updateScheduledActionCmd.Flags().String("start-time", "", "The start time in UTC of the scheduled action to update to.")
	redshiftServerless_updateScheduledActionCmd.Flags().String("target-action", "", "")
	redshiftServerless_updateScheduledActionCmd.Flag("no-enabled").Hidden = true
	redshiftServerless_updateScheduledActionCmd.MarkFlagRequired("scheduled-action-name")
	redshiftServerlessCmd.AddCommand(redshiftServerless_updateScheduledActionCmd)
}
