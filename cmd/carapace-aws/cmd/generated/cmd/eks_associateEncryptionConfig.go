package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var eks_associateEncryptionConfigCmd = &cobra.Command{
	Use:   "associate-encryption-config",
	Short: "Associates an encryption configuration to an existing cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(eks_associateEncryptionConfigCmd).Standalone()

	eks_associateEncryptionConfigCmd.Flags().String("client-request-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	eks_associateEncryptionConfigCmd.Flags().String("cluster-name", "", "The name of your cluster.")
	eks_associateEncryptionConfigCmd.Flags().String("encryption-config", "", "The configuration you are using for encryption.")
	eks_associateEncryptionConfigCmd.MarkFlagRequired("cluster-name")
	eks_associateEncryptionConfigCmd.MarkFlagRequired("encryption-config")
	eksCmd.AddCommand(eks_associateEncryptionConfigCmd)
}
