package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mgn_updateSourceServerReplicationTypeCmd = &cobra.Command{
	Use:   "update-source-server-replication-type",
	Short: "Allows you to change between the AGENT\\_BASED replication type and the SNAPSHOT\\_SHIPPING replication type.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mgn_updateSourceServerReplicationTypeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mgn_updateSourceServerReplicationTypeCmd).Standalone()

		mgn_updateSourceServerReplicationTypeCmd.Flags().String("account-id", "", "Account ID on which to update replication type.")
		mgn_updateSourceServerReplicationTypeCmd.Flags().String("replication-type", "", "Replication type to which to update source server.")
		mgn_updateSourceServerReplicationTypeCmd.Flags().String("source-server-id", "", "ID of source server on which to update replication type.")
		mgn_updateSourceServerReplicationTypeCmd.MarkFlagRequired("replication-type")
		mgn_updateSourceServerReplicationTypeCmd.MarkFlagRequired("source-server-id")
	})
	mgnCmd.AddCommand(mgn_updateSourceServerReplicationTypeCmd)
}
