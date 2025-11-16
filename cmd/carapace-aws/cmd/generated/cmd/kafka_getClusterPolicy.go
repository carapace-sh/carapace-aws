package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kafka_getClusterPolicyCmd = &cobra.Command{
	Use:   "get-cluster-policy",
	Short: "Get the MSK cluster policy specified by the Amazon Resource Name (ARN) in the request.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kafka_getClusterPolicyCmd).Standalone()

	kafka_getClusterPolicyCmd.Flags().String("cluster-arn", "", "The Amazon Resource Name (ARN) of the cluster.")
	kafka_getClusterPolicyCmd.MarkFlagRequired("cluster-arn")
	kafkaCmd.AddCommand(kafka_getClusterPolicyCmd)
}
