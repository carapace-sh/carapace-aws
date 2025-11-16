package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_modifyScheduledActionCmd = &cobra.Command{
	Use:   "modify-scheduled-action",
	Short: "Modifies a scheduled action.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_modifyScheduledActionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(redshift_modifyScheduledActionCmd).Standalone()

		redshift_modifyScheduledActionCmd.Flags().String("enable", "", "A modified enable flag of the scheduled action.")
		redshift_modifyScheduledActionCmd.Flags().String("end-time", "", "A modified end time of the scheduled action.")
		redshift_modifyScheduledActionCmd.Flags().String("iam-role", "", "A different IAM role to assume to run the target action.")
		redshift_modifyScheduledActionCmd.Flags().String("schedule", "", "A modified schedule in either `at( )` or `cron( )` format.")
		redshift_modifyScheduledActionCmd.Flags().String("scheduled-action-description", "", "A modified description of the scheduled action.")
		redshift_modifyScheduledActionCmd.Flags().String("scheduled-action-name", "", "The name of the scheduled action to modify.")
		redshift_modifyScheduledActionCmd.Flags().String("start-time", "", "A modified start time of the scheduled action.")
		redshift_modifyScheduledActionCmd.Flags().String("target-action", "", "A modified JSON format of the scheduled action.")
		redshift_modifyScheduledActionCmd.MarkFlagRequired("scheduled-action-name")
	})
	redshiftCmd.AddCommand(redshift_modifyScheduledActionCmd)
}
