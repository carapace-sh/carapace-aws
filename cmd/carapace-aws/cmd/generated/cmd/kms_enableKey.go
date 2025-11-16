package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kms_enableKeyCmd = &cobra.Command{
	Use:   "enable-key",
	Short: "Sets the key state of a KMS key to enabled.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kms_enableKeyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kms_enableKeyCmd).Standalone()

		kms_enableKeyCmd.Flags().String("key-id", "", "Identifies the KMS key to enable.")
		kms_enableKeyCmd.MarkFlagRequired("key-id")
	})
	kmsCmd.AddCommand(kms_enableKeyCmd)
}
