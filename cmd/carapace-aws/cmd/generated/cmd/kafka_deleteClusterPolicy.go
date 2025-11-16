package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kafka_deleteClusterPolicyCmd = &cobra.Command{
	Use:   "delete-cluster-policy",
	Short: "Deletes the MSK cluster policy specified by the Amazon Resource Name (ARN) in the request.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kafka_deleteClusterPolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kafka_deleteClusterPolicyCmd).Standalone()

		kafka_deleteClusterPolicyCmd.Flags().String("cluster-arn", "", "The Amazon Resource Name (ARN) of the cluster.")
		kafka_deleteClusterPolicyCmd.MarkFlagRequired("cluster-arn")
	})
	kafkaCmd.AddCommand(kafka_deleteClusterPolicyCmd)
}
