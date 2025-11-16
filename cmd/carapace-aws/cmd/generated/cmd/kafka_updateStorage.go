package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kafka_updateStorageCmd = &cobra.Command{
	Use:   "update-storage",
	Short: "Updates cluster broker volume size (or) sets cluster storage mode to TIERED.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kafka_updateStorageCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kafka_updateStorageCmd).Standalone()

		kafka_updateStorageCmd.Flags().String("cluster-arn", "", "The Amazon Resource Name (ARN) of the cluster to be updated.")
		kafka_updateStorageCmd.Flags().String("current-version", "", "The version of cluster to update from.")
		kafka_updateStorageCmd.Flags().String("provisioned-throughput", "", "EBS volume provisioned throughput information.")
		kafka_updateStorageCmd.Flags().String("storage-mode", "", "Controls storage mode for supported storage tiers.")
		kafka_updateStorageCmd.Flags().String("volume-size-gb", "", "size of the EBS volume to update.")
		kafka_updateStorageCmd.MarkFlagRequired("cluster-arn")
		kafka_updateStorageCmd.MarkFlagRequired("current-version")
	})
	kafkaCmd.AddCommand(kafka_updateStorageCmd)
}
