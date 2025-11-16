package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kafka_describeVpcConnectionCmd = &cobra.Command{
	Use:   "describe-vpc-connection",
	Short: "Returns a description of this MSK VPC connection.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kafka_describeVpcConnectionCmd).Standalone()

	kafka_describeVpcConnectionCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) that uniquely identifies a MSK VPC connection.")
	kafka_describeVpcConnectionCmd.MarkFlagRequired("arn")
	kafkaCmd.AddCommand(kafka_describeVpcConnectionCmd)
}
