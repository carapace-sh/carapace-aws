package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kafka_updateMonitoringCmd = &cobra.Command{
	Use:   "update-monitoring",
	Short: "Updates the monitoring settings for the cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kafka_updateMonitoringCmd).Standalone()

	kafka_updateMonitoringCmd.Flags().String("cluster-arn", "", "The Amazon Resource Name (ARN) that uniquely identifies the cluster.")
	kafka_updateMonitoringCmd.Flags().String("current-version", "", "The version of the MSK cluster to update.")
	kafka_updateMonitoringCmd.Flags().String("enhanced-monitoring", "", "Specifies which Apache Kafka metrics Amazon MSK gathers and sends to Amazon CloudWatch for this cluster.")
	kafka_updateMonitoringCmd.Flags().String("logging-info", "", "")
	kafka_updateMonitoringCmd.Flags().String("open-monitoring", "", "The settings for open monitoring.")
	kafka_updateMonitoringCmd.MarkFlagRequired("cluster-arn")
	kafka_updateMonitoringCmd.MarkFlagRequired("current-version")
	kafkaCmd.AddCommand(kafka_updateMonitoringCmd)
}
