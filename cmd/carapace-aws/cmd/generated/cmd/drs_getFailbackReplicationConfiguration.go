package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var drs_getFailbackReplicationConfigurationCmd = &cobra.Command{
	Use:   "get-failback-replication-configuration",
	Short: "Lists all Failback ReplicationConfigurations, filtered by Recovery Instance ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(drs_getFailbackReplicationConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(drs_getFailbackReplicationConfigurationCmd).Standalone()

		drs_getFailbackReplicationConfigurationCmd.Flags().String("recovery-instance-id", "", "The ID of the Recovery Instance whose failback replication configuration should be returned.")
		drs_getFailbackReplicationConfigurationCmd.MarkFlagRequired("recovery-instance-id")
	})
	drsCmd.AddCommand(drs_getFailbackReplicationConfigurationCmd)
}
