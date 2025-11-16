package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mgn_removeSourceServerActionCmd = &cobra.Command{
	Use:   "remove-source-server-action",
	Short: "Remove source server post migration custom action.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mgn_removeSourceServerActionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mgn_removeSourceServerActionCmd).Standalone()

		mgn_removeSourceServerActionCmd.Flags().String("account-id", "", "Source server post migration account ID.")
		mgn_removeSourceServerActionCmd.Flags().String("action-id", "", "Source server post migration custom action ID to remove.")
		mgn_removeSourceServerActionCmd.Flags().String("source-server-id", "", "Source server ID of the post migration custom action to remove.")
		mgn_removeSourceServerActionCmd.MarkFlagRequired("action-id")
		mgn_removeSourceServerActionCmd.MarkFlagRequired("source-server-id")
	})
	mgnCmd.AddCommand(mgn_removeSourceServerActionCmd)
}
