package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deadline_updateMonitorCmd = &cobra.Command{
	Use:   "update-monitor",
	Short: "Modifies the settings for a Deadline Cloud monitor.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deadline_updateMonitorCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(deadline_updateMonitorCmd).Standalone()

		deadline_updateMonitorCmd.Flags().String("display-name", "", "The new value to use for the monitor's display name.")
		deadline_updateMonitorCmd.Flags().String("monitor-id", "", "The unique identifier of the monitor to update.")
		deadline_updateMonitorCmd.Flags().String("role-arn", "", "The Amazon Resource Name (ARN) of the new IAM role to use with the monitor.")
		deadline_updateMonitorCmd.Flags().String("subdomain", "", "The new value of the subdomain to use when forming the monitor URL.")
		deadline_updateMonitorCmd.MarkFlagRequired("monitor-id")
	})
	deadlineCmd.AddCommand(deadline_updateMonitorCmd)
}
