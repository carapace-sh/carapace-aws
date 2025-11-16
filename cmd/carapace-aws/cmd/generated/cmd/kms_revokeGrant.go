package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kms_revokeGrantCmd = &cobra.Command{
	Use:   "revoke-grant",
	Short: "Deletes the specified grant.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kms_revokeGrantCmd).Standalone()

	kms_revokeGrantCmd.Flags().String("dry-run", "", "Checks if your request will succeed.")
	kms_revokeGrantCmd.Flags().String("grant-id", "", "Identifies the grant to revoke.")
	kms_revokeGrantCmd.Flags().String("key-id", "", "A unique identifier for the KMS key associated with the grant.")
	kms_revokeGrantCmd.MarkFlagRequired("grant-id")
	kms_revokeGrantCmd.MarkFlagRequired("key-id")
	kmsCmd.AddCommand(kms_revokeGrantCmd)
}
