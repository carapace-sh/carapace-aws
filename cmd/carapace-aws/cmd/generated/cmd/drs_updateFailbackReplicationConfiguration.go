package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var drs_updateFailbackReplicationConfigurationCmd = &cobra.Command{
	Use:   "update-failback-replication-configuration",
	Short: "Allows you to update the failback replication configuration of a Recovery Instance by ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(drs_updateFailbackReplicationConfigurationCmd).Standalone()

	drs_updateFailbackReplicationConfigurationCmd.Flags().String("bandwidth-throttling", "", "Configure bandwidth throttling for the outbound data transfer rate of the Recovery Instance in Mbps.")
	drs_updateFailbackReplicationConfigurationCmd.Flags().String("name", "", "The name of the Failback Replication Configuration.")
	drs_updateFailbackReplicationConfigurationCmd.Flags().Bool("no-use-private-ip", false, "Whether to use Private IP for the failback replication of the Recovery Instance.")
	drs_updateFailbackReplicationConfigurationCmd.Flags().String("recovery-instance-id", "", "The ID of the Recovery Instance.")
	drs_updateFailbackReplicationConfigurationCmd.Flags().Bool("use-private-ip", false, "Whether to use Private IP for the failback replication of the Recovery Instance.")
	drs_updateFailbackReplicationConfigurationCmd.Flag("no-use-private-ip").Hidden = true
	drs_updateFailbackReplicationConfigurationCmd.MarkFlagRequired("recovery-instance-id")
	drsCmd.AddCommand(drs_updateFailbackReplicationConfigurationCmd)
}
