package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var efs_createReplicationConfigurationCmd = &cobra.Command{
	Use:   "create-replication-configuration",
	Short: "Creates a replication conﬁguration to either a new or existing EFS file system.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(efs_createReplicationConfigurationCmd).Standalone()

	efs_createReplicationConfigurationCmd.Flags().String("destinations", "", "An array of destination configuration objects.")
	efs_createReplicationConfigurationCmd.Flags().String("source-file-system-id", "", "Specifies the Amazon EFS file system that you want to replicate.")
	efs_createReplicationConfigurationCmd.MarkFlagRequired("destinations")
	efs_createReplicationConfigurationCmd.MarkFlagRequired("source-file-system-id")
	efsCmd.AddCommand(efs_createReplicationConfigurationCmd)
}
