package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kafka_deleteReplicatorCmd = &cobra.Command{
	Use:   "delete-replicator",
	Short: "Deletes a replicator.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kafka_deleteReplicatorCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kafka_deleteReplicatorCmd).Standalone()

		kafka_deleteReplicatorCmd.Flags().String("current-version", "", "The current version of the replicator.")
		kafka_deleteReplicatorCmd.Flags().String("replicator-arn", "", "The Amazon Resource Name (ARN) of the replicator to be deleted.")
		kafka_deleteReplicatorCmd.MarkFlagRequired("replicator-arn")
	})
	kafkaCmd.AddCommand(kafka_deleteReplicatorCmd)
}
