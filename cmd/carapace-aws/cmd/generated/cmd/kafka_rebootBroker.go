package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kafka_rebootBrokerCmd = &cobra.Command{
	Use:   "reboot-broker",
	Short: "Reboots brokers.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kafka_rebootBrokerCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kafka_rebootBrokerCmd).Standalone()

		kafka_rebootBrokerCmd.Flags().String("broker-ids", "", "The list of broker IDs to be rebooted.")
		kafka_rebootBrokerCmd.Flags().String("cluster-arn", "", "The Amazon Resource Name (ARN) of the cluster to be updated.")
		kafka_rebootBrokerCmd.MarkFlagRequired("broker-ids")
		kafka_rebootBrokerCmd.MarkFlagRequired("cluster-arn")
	})
	kafkaCmd.AddCommand(kafka_rebootBrokerCmd)
}
