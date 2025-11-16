package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mgn_getReplicationConfigurationCmd = &cobra.Command{
	Use:   "get-replication-configuration",
	Short: "Lists all ReplicationConfigurations, filtered by Source Server ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mgn_getReplicationConfigurationCmd).Standalone()

	mgn_getReplicationConfigurationCmd.Flags().String("account-id", "", "Request to get Replication Configuration by Account ID.")
	mgn_getReplicationConfigurationCmd.Flags().String("source-server-id", "", "Request to get Replication Configuration by Source Server ID.")
	mgn_getReplicationConfigurationCmd.MarkFlagRequired("source-server-id")
	mgnCmd.AddCommand(mgn_getReplicationConfigurationCmd)
}
