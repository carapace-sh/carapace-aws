package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var auditmanager_listNotificationsCmd = &cobra.Command{
	Use:   "list-notifications",
	Short: "Returns a list of all Audit Manager notifications.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auditmanager_listNotificationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(auditmanager_listNotificationsCmd).Standalone()

		auditmanager_listNotificationsCmd.Flags().String("max-results", "", "Represents the maximum number of results on a page or for an API request call.")
		auditmanager_listNotificationsCmd.Flags().String("next-token", "", "The pagination token that's used to fetch the next set of results.")
	})
	auditmanagerCmd.AddCommand(auditmanager_listNotificationsCmd)
}
