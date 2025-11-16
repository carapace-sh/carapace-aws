package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mpa_deleteIdentitySourceCmd = &cobra.Command{
	Use:   "delete-identity-source",
	Short: "Deletes an identity source.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mpa_deleteIdentitySourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mpa_deleteIdentitySourceCmd).Standalone()

		mpa_deleteIdentitySourceCmd.Flags().String("identity-source-arn", "", "Amazon Resource Name (ARN) for identity source.")
		mpa_deleteIdentitySourceCmd.MarkFlagRequired("identity-source-arn")
	})
	mpaCmd.AddCommand(mpa_deleteIdentitySourceCmd)
}
