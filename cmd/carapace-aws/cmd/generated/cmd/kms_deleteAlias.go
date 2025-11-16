package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kms_deleteAliasCmd = &cobra.Command{
	Use:   "delete-alias",
	Short: "Deletes the specified alias.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kms_deleteAliasCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kms_deleteAliasCmd).Standalone()

		kms_deleteAliasCmd.Flags().String("alias-name", "", "The alias to be deleted.")
		kms_deleteAliasCmd.MarkFlagRequired("alias-name")
	})
	kmsCmd.AddCommand(kms_deleteAliasCmd)
}
