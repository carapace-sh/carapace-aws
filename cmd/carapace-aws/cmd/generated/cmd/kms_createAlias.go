package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kms_createAliasCmd = &cobra.Command{
	Use:   "create-alias",
	Short: "Creates a friendly name for a KMS key.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kms_createAliasCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kms_createAliasCmd).Standalone()

		kms_createAliasCmd.Flags().String("alias-name", "", "Specifies the alias name.")
		kms_createAliasCmd.Flags().String("target-key-id", "", "Associates the alias with the specified [customer managed key](https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#customer-mgn-key).")
		kms_createAliasCmd.MarkFlagRequired("alias-name")
		kms_createAliasCmd.MarkFlagRequired("target-key-id")
	})
	kmsCmd.AddCommand(kms_createAliasCmd)
}
