package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kafka_batchDisassociateScramSecretCmd = &cobra.Command{
	Use:   "batch-disassociate-scram-secret",
	Short: "Disassociates one or more Scram Secrets from an Amazon MSK cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kafka_batchDisassociateScramSecretCmd).Standalone()

	kafka_batchDisassociateScramSecretCmd.Flags().String("cluster-arn", "", "The Amazon Resource Name (ARN) of the cluster to be updated.")
	kafka_batchDisassociateScramSecretCmd.Flags().String("secret-arn-list", "", "List of AWS Secrets Manager secret ARNs.")
	kafka_batchDisassociateScramSecretCmd.MarkFlagRequired("cluster-arn")
	kafka_batchDisassociateScramSecretCmd.MarkFlagRequired("secret-arn-list")
	kafkaCmd.AddCommand(kafka_batchDisassociateScramSecretCmd)
}
