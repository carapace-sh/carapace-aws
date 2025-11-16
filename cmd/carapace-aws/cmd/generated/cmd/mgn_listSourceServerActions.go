package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mgn_listSourceServerActionsCmd = &cobra.Command{
	Use:   "list-source-server-actions",
	Short: "List source server post migration custom actions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mgn_listSourceServerActionsCmd).Standalone()

	mgn_listSourceServerActionsCmd.Flags().String("account-id", "", "Account ID to return when listing source server post migration custom actions.")
	mgn_listSourceServerActionsCmd.Flags().String("filters", "", "Filters to apply when listing source server post migration custom actions.")
	mgn_listSourceServerActionsCmd.Flags().String("max-results", "", "Maximum amount of items to return when listing source server post migration custom actions.")
	mgn_listSourceServerActionsCmd.Flags().String("next-token", "", "Next token to use when listing source server post migration custom actions.")
	mgn_listSourceServerActionsCmd.Flags().String("source-server-id", "", "Source server ID.")
	mgn_listSourceServerActionsCmd.MarkFlagRequired("source-server-id")
	mgnCmd.AddCommand(mgn_listSourceServerActionsCmd)
}
