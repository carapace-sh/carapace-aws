package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kafka_updateBrokerTypeCmd = &cobra.Command{
	Use:   "update-broker-type",
	Short: "Updates EC2 instance type.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kafka_updateBrokerTypeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kafka_updateBrokerTypeCmd).Standalone()

		kafka_updateBrokerTypeCmd.Flags().String("cluster-arn", "", "The Amazon Resource Name (ARN) that uniquely identifies the cluster.")
		kafka_updateBrokerTypeCmd.Flags().String("current-version", "", "The cluster version that you want to change.")
		kafka_updateBrokerTypeCmd.Flags().String("target-instance-type", "", "The Amazon MSK broker type that you want all of the brokers in this cluster to be.")
		kafka_updateBrokerTypeCmd.MarkFlagRequired("cluster-arn")
		kafka_updateBrokerTypeCmd.MarkFlagRequired("current-version")
		kafka_updateBrokerTypeCmd.MarkFlagRequired("target-instance-type")
	})
	kafkaCmd.AddCommand(kafka_updateBrokerTypeCmd)
}
