package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kms_updateKeyDescriptionCmd = &cobra.Command{
	Use:   "update-key-description",
	Short: "Updates the description of a KMS key.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kms_updateKeyDescriptionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kms_updateKeyDescriptionCmd).Standalone()

		kms_updateKeyDescriptionCmd.Flags().String("description", "", "New description for the KMS key.")
		kms_updateKeyDescriptionCmd.Flags().String("key-id", "", "Updates the description of the specified KMS key.")
		kms_updateKeyDescriptionCmd.MarkFlagRequired("description")
		kms_updateKeyDescriptionCmd.MarkFlagRequired("key-id")
	})
	kmsCmd.AddCommand(kms_updateKeyDescriptionCmd)
}
