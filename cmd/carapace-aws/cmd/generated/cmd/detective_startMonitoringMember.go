package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var detective_startMonitoringMemberCmd = &cobra.Command{
	Use:   "start-monitoring-member",
	Short: "Sends a request to enable data ingest for a member account that has a status of `ACCEPTED_BUT_DISABLED`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(detective_startMonitoringMemberCmd).Standalone()

	detective_startMonitoringMemberCmd.Flags().String("account-id", "", "The account ID of the member account to try to enable.")
	detective_startMonitoringMemberCmd.Flags().String("graph-arn", "", "The ARN of the behavior graph.")
	detective_startMonitoringMemberCmd.MarkFlagRequired("account-id")
	detective_startMonitoringMemberCmd.MarkFlagRequired("graph-arn")
	detectiveCmd.AddCommand(detective_startMonitoringMemberCmd)
}
