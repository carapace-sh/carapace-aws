package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var opensearch_updateScheduledActionCmd = &cobra.Command{
	Use:   "update-scheduled-action",
	Short: "Reschedules a planned domain configuration change for a later time.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opensearch_updateScheduledActionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(opensearch_updateScheduledActionCmd).Standalone()

		opensearch_updateScheduledActionCmd.Flags().String("action-id", "", "The unique identifier of the action to reschedule.")
		opensearch_updateScheduledActionCmd.Flags().String("action-type", "", "The type of action to reschedule.")
		opensearch_updateScheduledActionCmd.Flags().String("desired-start-time", "", "The time to implement the change, in Coordinated Universal Time (UTC).")
		opensearch_updateScheduledActionCmd.Flags().String("domain-name", "", "The name of the domain to reschedule an action for.")
		opensearch_updateScheduledActionCmd.Flags().String("schedule-at", "", "When to schedule the action.")
		opensearch_updateScheduledActionCmd.MarkFlagRequired("action-id")
		opensearch_updateScheduledActionCmd.MarkFlagRequired("action-type")
		opensearch_updateScheduledActionCmd.MarkFlagRequired("domain-name")
		opensearch_updateScheduledActionCmd.MarkFlagRequired("schedule-at")
	})
	opensearchCmd.AddCommand(opensearch_updateScheduledActionCmd)
}
