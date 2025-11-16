package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kafka_listKafkaVersionsCmd = &cobra.Command{
	Use:   "list-kafka-versions",
	Short: "Returns a list of Apache Kafka versions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kafka_listKafkaVersionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kafka_listKafkaVersionsCmd).Standalone()

		kafka_listKafkaVersionsCmd.Flags().String("max-results", "", "The maximum number of results to return in the response.")
		kafka_listKafkaVersionsCmd.Flags().String("next-token", "", "The paginated results marker.")
	})
	kafkaCmd.AddCommand(kafka_listKafkaVersionsCmd)
}
