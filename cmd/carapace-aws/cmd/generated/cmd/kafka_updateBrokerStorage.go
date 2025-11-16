package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kafka_updateBrokerStorageCmd = &cobra.Command{
	Use:   "update-broker-storage",
	Short: "Updates the EBS storage associated with MSK brokers.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kafka_updateBrokerStorageCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kafka_updateBrokerStorageCmd).Standalone()

		kafka_updateBrokerStorageCmd.Flags().String("cluster-arn", "", "The Amazon Resource Name (ARN) that uniquely identifies the cluster.")
		kafka_updateBrokerStorageCmd.Flags().String("current-version", "", "The version of cluster to update from.")
		kafka_updateBrokerStorageCmd.Flags().String("target-broker-ebsvolume-info", "", "Describes the target volume size and the ID of the broker to apply the update to.")
		kafka_updateBrokerStorageCmd.MarkFlagRequired("cluster-arn")
		kafka_updateBrokerStorageCmd.MarkFlagRequired("current-version")
		kafka_updateBrokerStorageCmd.MarkFlagRequired("target-broker-ebsvolume-info")
	})
	kafkaCmd.AddCommand(kafka_updateBrokerStorageCmd)
}
