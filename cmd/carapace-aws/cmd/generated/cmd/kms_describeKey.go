package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kms_describeKeyCmd = &cobra.Command{
	Use:   "describe-key",
	Short: "Provides detailed information about a KMS key.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kms_describeKeyCmd).Standalone()

	kms_describeKeyCmd.Flags().String("grant-tokens", "", "A list of grant tokens.")
	kms_describeKeyCmd.Flags().String("key-id", "", "Describes the specified KMS key.")
	kms_describeKeyCmd.MarkFlagRequired("key-id")
	kmsCmd.AddCommand(kms_describeKeyCmd)
}
