package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kafka_updateRebalancingCmd = &cobra.Command{
	Use:   "update-rebalancing",
	Short: "Use this resource to update the intelligent rebalancing status of an Amazon MSK Provisioned cluster with Express brokers.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kafka_updateRebalancingCmd).Standalone()

	kafka_updateRebalancingCmd.Flags().String("cluster-arn", "", "The Amazon Resource Name (ARN) of the cluster.")
	kafka_updateRebalancingCmd.Flags().String("current-version", "", "The current version of the cluster.")
	kafka_updateRebalancingCmd.Flags().String("rebalancing", "", "Specifies if intelligent rebalancing should be turned on for your cluster.")
	kafka_updateRebalancingCmd.MarkFlagRequired("cluster-arn")
	kafka_updateRebalancingCmd.MarkFlagRequired("current-version")
	kafka_updateRebalancingCmd.MarkFlagRequired("rebalancing")
	kafkaCmd.AddCommand(kafka_updateRebalancingCmd)
}
