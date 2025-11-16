package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kafka_putClusterPolicyCmd = &cobra.Command{
	Use:   "put-cluster-policy",
	Short: "Creates or updates the MSK cluster policy specified by the cluster Amazon Resource Name (ARN) in the request.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kafka_putClusterPolicyCmd).Standalone()

	kafka_putClusterPolicyCmd.Flags().String("cluster-arn", "", "The Amazon Resource Name (ARN) of the cluster.")
	kafka_putClusterPolicyCmd.Flags().String("current-version", "", "The policy version.")
	kafka_putClusterPolicyCmd.Flags().String("policy", "", "The policy.")
	kafka_putClusterPolicyCmd.MarkFlagRequired("cluster-arn")
	kafka_putClusterPolicyCmd.MarkFlagRequired("policy")
	kafkaCmd.AddCommand(kafka_putClusterPolicyCmd)
}
