package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kafka_createClusterV2Cmd = &cobra.Command{
	Use:   "create-cluster-v2",
	Short: "Creates a new MSK cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kafka_createClusterV2Cmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kafka_createClusterV2Cmd).Standalone()

		kafka_createClusterV2Cmd.Flags().String("cluster-name", "", "The name of the cluster.")
		kafka_createClusterV2Cmd.Flags().String("provisioned", "", "Information about the provisioned cluster.")
		kafka_createClusterV2Cmd.Flags().String("serverless", "", "Information about the serverless cluster.")
		kafka_createClusterV2Cmd.Flags().String("tags", "", "A map of tags that you want the cluster to have.")
		kafka_createClusterV2Cmd.MarkFlagRequired("cluster-name")
	})
	kafkaCmd.AddCommand(kafka_createClusterV2Cmd)
}
