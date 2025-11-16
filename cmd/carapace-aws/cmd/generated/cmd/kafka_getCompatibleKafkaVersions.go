package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kafka_getCompatibleKafkaVersionsCmd = &cobra.Command{
	Use:   "get-compatible-kafka-versions",
	Short: "Gets the Apache Kafka versions to which you can update the MSK cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kafka_getCompatibleKafkaVersionsCmd).Standalone()

	kafka_getCompatibleKafkaVersionsCmd.Flags().String("cluster-arn", "", "The Amazon Resource Name (ARN) of the cluster check.")
	kafkaCmd.AddCommand(kafka_getCompatibleKafkaVersionsCmd)
}
