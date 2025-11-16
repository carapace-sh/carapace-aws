package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kafka_batchAssociateScramSecretCmd = &cobra.Command{
	Use:   "batch-associate-scram-secret",
	Short: "Associates one or more Scram Secrets with an Amazon MSK cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kafka_batchAssociateScramSecretCmd).Standalone()

	kafka_batchAssociateScramSecretCmd.Flags().String("cluster-arn", "", "The Amazon Resource Name (ARN) of the cluster to be updated.")
	kafka_batchAssociateScramSecretCmd.Flags().String("secret-arn-list", "", "List of AWS Secrets Manager secret ARNs.")
	kafka_batchAssociateScramSecretCmd.MarkFlagRequired("cluster-arn")
	kafka_batchAssociateScramSecretCmd.MarkFlagRequired("secret-arn-list")
	kafkaCmd.AddCommand(kafka_batchAssociateScramSecretCmd)
}
