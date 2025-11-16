package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codecommit_updateRepositoryEncryptionKeyCmd = &cobra.Command{
	Use:   "update-repository-encryption-key",
	Short: "Updates the Key Management Service encryption key used to encrypt and decrypt a CodeCommit repository.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codecommit_updateRepositoryEncryptionKeyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codecommit_updateRepositoryEncryptionKeyCmd).Standalone()

		codecommit_updateRepositoryEncryptionKeyCmd.Flags().String("kms-key-id", "", "The ID of the encryption key.")
		codecommit_updateRepositoryEncryptionKeyCmd.Flags().String("repository-name", "", "The name of the repository for which you want to update the KMS encryption key used to encrypt and decrypt the repository.")
		codecommit_updateRepositoryEncryptionKeyCmd.MarkFlagRequired("kms-key-id")
		codecommit_updateRepositoryEncryptionKeyCmd.MarkFlagRequired("repository-name")
	})
	codecommitCmd.AddCommand(codecommit_updateRepositoryEncryptionKeyCmd)
}
