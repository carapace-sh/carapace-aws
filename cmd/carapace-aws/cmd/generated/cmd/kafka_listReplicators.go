package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kafka_listReplicatorsCmd = &cobra.Command{
	Use:   "list-replicators",
	Short: "Lists the replicators.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kafka_listReplicatorsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kafka_listReplicatorsCmd).Standalone()

		kafka_listReplicatorsCmd.Flags().String("max-results", "", "The maximum number of results to return in the response.")
		kafka_listReplicatorsCmd.Flags().String("next-token", "", "If the response of ListReplicators is truncated, it returns a NextToken in the response.")
		kafka_listReplicatorsCmd.Flags().String("replicator-name-filter", "", "Returns replicators starting with given name.")
	})
	kafkaCmd.AddCommand(kafka_listReplicatorsCmd)
}
