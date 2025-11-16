package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kafka_describeReplicatorCmd = &cobra.Command{
	Use:   "describe-replicator",
	Short: "Describes a replicator.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kafka_describeReplicatorCmd).Standalone()

	kafka_describeReplicatorCmd.Flags().String("replicator-arn", "", "The Amazon Resource Name (ARN) of the replicator to be described.")
	kafka_describeReplicatorCmd.MarkFlagRequired("replicator-arn")
	kafkaCmd.AddCommand(kafka_describeReplicatorCmd)
}
