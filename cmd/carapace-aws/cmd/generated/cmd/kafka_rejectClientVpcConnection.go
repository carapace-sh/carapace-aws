package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kafka_rejectClientVpcConnectionCmd = &cobra.Command{
	Use:   "reject-client-vpc-connection",
	Short: "Returns empty response.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kafka_rejectClientVpcConnectionCmd).Standalone()

	kafka_rejectClientVpcConnectionCmd.Flags().String("cluster-arn", "", "The Amazon Resource Name (ARN) of the cluster.")
	kafka_rejectClientVpcConnectionCmd.Flags().String("vpc-connection-arn", "", "The VPC connection ARN.")
	kafka_rejectClientVpcConnectionCmd.MarkFlagRequired("cluster-arn")
	kafka_rejectClientVpcConnectionCmd.MarkFlagRequired("vpc-connection-arn")
	kafkaCmd.AddCommand(kafka_rejectClientVpcConnectionCmd)
}
