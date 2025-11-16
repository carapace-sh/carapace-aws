package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var secretsmanager_removeRegionsFromReplicationCmd = &cobra.Command{
	Use:   "remove-regions-from-replication",
	Short: "For a secret that is replicated to other Regions, deletes the secret replicas from the Regions you specify.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(secretsmanager_removeRegionsFromReplicationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(secretsmanager_removeRegionsFromReplicationCmd).Standalone()

		secretsmanager_removeRegionsFromReplicationCmd.Flags().String("remove-replica-regions", "", "The Regions of the replicas to remove.")
		secretsmanager_removeRegionsFromReplicationCmd.Flags().String("secret-id", "", "The ARN or name of the secret.")
		secretsmanager_removeRegionsFromReplicationCmd.MarkFlagRequired("remove-replica-regions")
		secretsmanager_removeRegionsFromReplicationCmd.MarkFlagRequired("secret-id")
	})
	secretsmanagerCmd.AddCommand(secretsmanager_removeRegionsFromReplicationCmd)
}
