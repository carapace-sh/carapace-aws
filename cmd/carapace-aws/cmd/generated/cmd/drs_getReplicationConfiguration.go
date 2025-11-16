package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var drs_getReplicationConfigurationCmd = &cobra.Command{
	Use:   "get-replication-configuration",
	Short: "Gets a ReplicationConfiguration, filtered by Source Server ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(drs_getReplicationConfigurationCmd).Standalone()

	drs_getReplicationConfigurationCmd.Flags().String("source-server-id", "", "The ID of the Source Serve for this Replication Configuration.r")
	drs_getReplicationConfigurationCmd.MarkFlagRequired("source-server-id")
	drsCmd.AddCommand(drs_getReplicationConfigurationCmd)
}
