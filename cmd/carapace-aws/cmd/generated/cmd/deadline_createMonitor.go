package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deadline_createMonitorCmd = &cobra.Command{
	Use:   "create-monitor",
	Short: "Creates an Amazon Web Services Deadline Cloud monitor that you can use to view your farms, queues, and fleets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deadline_createMonitorCmd).Standalone()

	deadline_createMonitorCmd.Flags().String("client-token", "", "The unique token which the server uses to recognize retries of the same request.")
	deadline_createMonitorCmd.Flags().String("display-name", "", "The name that you give the monitor that is displayed in the Deadline Cloud console.")
	deadline_createMonitorCmd.Flags().String("identity-center-instance-arn", "", "The Amazon Resource Name (ARN) of the IAM Identity Center instance that authenticates monitor users.")
	deadline_createMonitorCmd.Flags().String("role-arn", "", "The Amazon Resource Name (ARN) of the IAM role that the monitor uses to connect to Deadline Cloud.")
	deadline_createMonitorCmd.Flags().String("subdomain", "", "The subdomain to use when creating the monitor URL.")
	deadline_createMonitorCmd.Flags().String("tags", "", "The tags to add to your monitor.")
	deadline_createMonitorCmd.MarkFlagRequired("display-name")
	deadline_createMonitorCmd.MarkFlagRequired("identity-center-instance-arn")
	deadline_createMonitorCmd.MarkFlagRequired("role-arn")
	deadline_createMonitorCmd.MarkFlagRequired("subdomain")
	deadlineCmd.AddCommand(deadline_createMonitorCmd)
}
