package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mgn_pauseReplicationCmd = &cobra.Command{
	Use:   "pause-replication",
	Short: "Pause Replication.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mgn_pauseReplicationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mgn_pauseReplicationCmd).Standalone()

		mgn_pauseReplicationCmd.Flags().String("account-id", "", "Pause Replication Request account ID.")
		mgn_pauseReplicationCmd.Flags().String("source-server-id", "", "Pause Replication Request source server ID.")
		mgn_pauseReplicationCmd.MarkFlagRequired("source-server-id")
	})
	mgnCmd.AddCommand(mgn_pauseReplicationCmd)
}
