package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kms_updateAliasCmd = &cobra.Command{
	Use:   "update-alias",
	Short: "Associates an existing KMS alias with a different KMS key.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kms_updateAliasCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kms_updateAliasCmd).Standalone()

		kms_updateAliasCmd.Flags().String("alias-name", "", "Identifies the alias that is changing its KMS key.")
		kms_updateAliasCmd.Flags().String("target-key-id", "", "Identifies the [customer managed key](https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#customer-mgn-key) to associate with the alias.")
		kms_updateAliasCmd.MarkFlagRequired("alias-name")
		kms_updateAliasCmd.MarkFlagRequired("target-key-id")
	})
	kmsCmd.AddCommand(kms_updateAliasCmd)
}
