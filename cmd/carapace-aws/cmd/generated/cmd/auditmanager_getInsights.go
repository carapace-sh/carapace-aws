package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var auditmanager_getInsightsCmd = &cobra.Command{
	Use:   "get-insights",
	Short: "Gets the latest analytics data for all your current active assessments.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auditmanager_getInsightsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(auditmanager_getInsightsCmd).Standalone()

	})
	auditmanagerCmd.AddCommand(auditmanager_getInsightsCmd)
}
