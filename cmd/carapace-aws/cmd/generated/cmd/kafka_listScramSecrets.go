package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kafka_listScramSecretsCmd = &cobra.Command{
	Use:   "list-scram-secrets",
	Short: "Returns a list of the Scram Secrets associated with an Amazon MSK cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kafka_listScramSecretsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kafka_listScramSecretsCmd).Standalone()

		kafka_listScramSecretsCmd.Flags().String("cluster-arn", "", "The arn of the cluster.")
		kafka_listScramSecretsCmd.Flags().String("max-results", "", "The maxResults of the query.")
		kafka_listScramSecretsCmd.Flags().String("next-token", "", "The nextToken of the query.")
		kafka_listScramSecretsCmd.MarkFlagRequired("cluster-arn")
	})
	kafkaCmd.AddCommand(kafka_listScramSecretsCmd)
}
