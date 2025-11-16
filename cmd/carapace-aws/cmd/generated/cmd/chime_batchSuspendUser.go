package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chime_batchSuspendUserCmd = &cobra.Command{
	Use:   "batch-suspend-user",
	Short: "Suspends up to 50 users from a `Team` or `EnterpriseLWA` Amazon Chime account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chime_batchSuspendUserCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chime_batchSuspendUserCmd).Standalone()

		chime_batchSuspendUserCmd.Flags().String("account-id", "", "The Amazon Chime account ID.")
		chime_batchSuspendUserCmd.Flags().String("user-id-list", "", "The request containing the user IDs to suspend.")
		chime_batchSuspendUserCmd.MarkFlagRequired("account-id")
		chime_batchSuspendUserCmd.MarkFlagRequired("user-id-list")
	})
	chimeCmd.AddCommand(chime_batchSuspendUserCmd)
}
