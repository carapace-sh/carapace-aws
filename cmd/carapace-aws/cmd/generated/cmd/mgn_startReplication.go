package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mgn_startReplicationCmd = &cobra.Command{
	Use:   "start-replication",
	Short: "Starts replication for SNAPSHOT\\_SHIPPING agents.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mgn_startReplicationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mgn_startReplicationCmd).Standalone()

		mgn_startReplicationCmd.Flags().String("account-id", "", "Account ID on which to start replication.")
		mgn_startReplicationCmd.Flags().String("source-server-id", "", "ID of source server on which to start replication.")
		mgn_startReplicationCmd.MarkFlagRequired("source-server-id")
	})
	mgnCmd.AddCommand(mgn_startReplicationCmd)
}
