package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kafka_deleteVpcConnectionCmd = &cobra.Command{
	Use:   "delete-vpc-connection",
	Short: "Deletes a MSK VPC connection.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kafka_deleteVpcConnectionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kafka_deleteVpcConnectionCmd).Standalone()

		kafka_deleteVpcConnectionCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) that uniquely identifies an MSK VPC connection.")
		kafka_deleteVpcConnectionCmd.MarkFlagRequired("arn")
	})
	kafkaCmd.AddCommand(kafka_deleteVpcConnectionCmd)
}
