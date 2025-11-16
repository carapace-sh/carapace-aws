package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecr_putReplicationConfigurationCmd = &cobra.Command{
	Use:   "put-replication-configuration",
	Short: "Creates or updates the replication configuration for a registry.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecr_putReplicationConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ecr_putReplicationConfigurationCmd).Standalone()

		ecr_putReplicationConfigurationCmd.Flags().String("replication-configuration", "", "An object representing the replication configuration for a registry.")
		ecr_putReplicationConfigurationCmd.MarkFlagRequired("replication-configuration")
	})
	ecrCmd.AddCommand(ecr_putReplicationConfigurationCmd)
}
