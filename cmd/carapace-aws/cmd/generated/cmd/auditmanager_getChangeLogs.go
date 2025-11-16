package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var auditmanager_getChangeLogsCmd = &cobra.Command{
	Use:   "get-change-logs",
	Short: "Gets a list of changelogs from Audit Manager.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auditmanager_getChangeLogsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(auditmanager_getChangeLogsCmd).Standalone()

		auditmanager_getChangeLogsCmd.Flags().String("assessment-id", "", "The unique identifier for the assessment.")
		auditmanager_getChangeLogsCmd.Flags().String("control-id", "", "The unique identifier for the control.")
		auditmanager_getChangeLogsCmd.Flags().String("control-set-id", "", "The unique identifier for the control set.")
		auditmanager_getChangeLogsCmd.Flags().String("max-results", "", "Represents the maximum number of results on a page or for an API request call.")
		auditmanager_getChangeLogsCmd.Flags().String("next-token", "", "The pagination token that's used to fetch the next set of results.")
		auditmanager_getChangeLogsCmd.MarkFlagRequired("assessment-id")
	})
	auditmanagerCmd.AddCommand(auditmanager_getChangeLogsCmd)
}
