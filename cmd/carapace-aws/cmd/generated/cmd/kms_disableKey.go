package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kms_disableKeyCmd = &cobra.Command{
	Use:   "disable-key",
	Short: "Sets the state of a KMS key to disabled.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kms_disableKeyCmd).Standalone()

	kms_disableKeyCmd.Flags().String("key-id", "", "Identifies the KMS key to disable.")
	kms_disableKeyCmd.MarkFlagRequired("key-id")
	kmsCmd.AddCommand(kms_disableKeyCmd)
}
