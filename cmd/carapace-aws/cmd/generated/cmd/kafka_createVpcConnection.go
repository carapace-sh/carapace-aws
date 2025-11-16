package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kafka_createVpcConnectionCmd = &cobra.Command{
	Use:   "create-vpc-connection",
	Short: "Creates a new MSK VPC connection.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kafka_createVpcConnectionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kafka_createVpcConnectionCmd).Standalone()

		kafka_createVpcConnectionCmd.Flags().String("authentication", "", "The authentication type of VPC connection.")
		kafka_createVpcConnectionCmd.Flags().String("client-subnets", "", "The list of client subnets.")
		kafka_createVpcConnectionCmd.Flags().String("security-groups", "", "The list of security groups.")
		kafka_createVpcConnectionCmd.Flags().String("tags", "", "A map of tags for the VPC connection.")
		kafka_createVpcConnectionCmd.Flags().String("target-cluster-arn", "", "The cluster Amazon Resource Name (ARN) for the VPC connection.")
		kafka_createVpcConnectionCmd.Flags().String("vpc-id", "", "The VPC ID of VPC connection.")
		kafka_createVpcConnectionCmd.MarkFlagRequired("authentication")
		kafka_createVpcConnectionCmd.MarkFlagRequired("client-subnets")
		kafka_createVpcConnectionCmd.MarkFlagRequired("security-groups")
		kafka_createVpcConnectionCmd.MarkFlagRequired("target-cluster-arn")
		kafka_createVpcConnectionCmd.MarkFlagRequired("vpc-id")
	})
	kafkaCmd.AddCommand(kafka_createVpcConnectionCmd)
}
