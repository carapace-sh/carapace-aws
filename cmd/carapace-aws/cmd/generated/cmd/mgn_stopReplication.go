package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mgn_stopReplicationCmd = &cobra.Command{
	Use:   "stop-replication",
	Short: "Stop Replication.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mgn_stopReplicationCmd).Standalone()

	mgn_stopReplicationCmd.Flags().String("account-id", "", "Stop Replication Request account ID.")
	mgn_stopReplicationCmd.Flags().String("source-server-id", "", "Stop Replication Request source server ID.")
	mgn_stopReplicationCmd.MarkFlagRequired("source-server-id")
	mgnCmd.AddCommand(mgn_stopReplicationCmd)
}
