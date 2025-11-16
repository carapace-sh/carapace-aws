package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kmsCmd = &cobra.Command{
	Use:   "kms",
	Short: "Key Management Service\n\nKey Management Service (KMS) is an encryption and key management web service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kmsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kmsCmd).Standalone()

	})
	rootCmd.AddCommand(kmsCmd)
}
