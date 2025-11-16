package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kafka_updateConnectivityCmd = &cobra.Command{
	Use:   "update-connectivity",
	Short: "Updates the cluster's connectivity configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kafka_updateConnectivityCmd).Standalone()

	kafka_updateConnectivityCmd.Flags().String("cluster-arn", "", "The Amazon Resource Name (ARN) of the configuration.")
	kafka_updateConnectivityCmd.Flags().String("connectivity-info", "", "Information about the broker access configuration.")
	kafka_updateConnectivityCmd.Flags().String("current-version", "", "The version of the MSK cluster to update.")
	kafka_updateConnectivityCmd.MarkFlagRequired("cluster-arn")
	kafka_updateConnectivityCmd.MarkFlagRequired("connectivity-info")
	kafka_updateConnectivityCmd.MarkFlagRequired("current-version")
	kafkaCmd.AddCommand(kafka_updateConnectivityCmd)
}
