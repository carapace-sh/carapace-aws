package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kms_retireGrantCmd = &cobra.Command{
	Use:   "retire-grant",
	Short: "Deletes a grant.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kms_retireGrantCmd).Standalone()

	kms_retireGrantCmd.Flags().String("dry-run", "", "Checks if your request will succeed.")
	kms_retireGrantCmd.Flags().String("grant-id", "", "Identifies the grant to retire.")
	kms_retireGrantCmd.Flags().String("grant-token", "", "Identifies the grant to be retired.")
	kms_retireGrantCmd.Flags().String("key-id", "", "The key ARN KMS key associated with the grant.")
	kmsCmd.AddCommand(kms_retireGrantCmd)
}
