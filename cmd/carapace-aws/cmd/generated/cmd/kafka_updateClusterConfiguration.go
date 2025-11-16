package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kafka_updateClusterConfigurationCmd = &cobra.Command{
	Use:   "update-cluster-configuration",
	Short: "Updates the cluster with the configuration that is specified in the request body.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kafka_updateClusterConfigurationCmd).Standalone()

	kafka_updateClusterConfigurationCmd.Flags().String("cluster-arn", "", "The Amazon Resource Name (ARN) that uniquely identifies the cluster.")
	kafka_updateClusterConfigurationCmd.Flags().String("configuration-info", "", "Represents the configuration that you want MSK to use for the brokers in a cluster.")
	kafka_updateClusterConfigurationCmd.Flags().String("current-version", "", "The version of the cluster that needs to be updated.")
	kafka_updateClusterConfigurationCmd.MarkFlagRequired("cluster-arn")
	kafka_updateClusterConfigurationCmd.MarkFlagRequired("configuration-info")
	kafka_updateClusterConfigurationCmd.MarkFlagRequired("current-version")
	kafkaCmd.AddCommand(kafka_updateClusterConfigurationCmd)
}
