package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var secretsmanager_stopReplicationToReplicaCmd = &cobra.Command{
	Use:   "stop-replication-to-replica",
	Short: "Removes the link between the replica secret and the primary secret and promotes the replica to a primary secret in the replica Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(secretsmanager_stopReplicationToReplicaCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(secretsmanager_stopReplicationToReplicaCmd).Standalone()

		secretsmanager_stopReplicationToReplicaCmd.Flags().String("secret-id", "", "The ARN of the primary secret.")
		secretsmanager_stopReplicationToReplicaCmd.MarkFlagRequired("secret-id")
	})
	secretsmanagerCmd.AddCommand(secretsmanager_stopReplicationToReplicaCmd)
}
