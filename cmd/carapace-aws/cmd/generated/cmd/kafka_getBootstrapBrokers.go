package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kafka_getBootstrapBrokersCmd = &cobra.Command{
	Use:   "get-bootstrap-brokers",
	Short: "A list of brokers that a client application can use to bootstrap.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kafka_getBootstrapBrokersCmd).Standalone()

	kafka_getBootstrapBrokersCmd.Flags().String("cluster-arn", "", "The Amazon Resource Name (ARN) that uniquely identifies the cluster.")
	kafka_getBootstrapBrokersCmd.MarkFlagRequired("cluster-arn")
	kafkaCmd.AddCommand(kafka_getBootstrapBrokersCmd)
}
