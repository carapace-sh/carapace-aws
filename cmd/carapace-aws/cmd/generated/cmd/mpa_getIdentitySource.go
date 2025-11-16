package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mpa_getIdentitySourceCmd = &cobra.Command{
	Use:   "get-identity-source",
	Short: "Returns details for an identity source.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mpa_getIdentitySourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mpa_getIdentitySourceCmd).Standalone()

		mpa_getIdentitySourceCmd.Flags().String("identity-source-arn", "", "Amazon Resource Name (ARN) for the identity source.")
		mpa_getIdentitySourceCmd.MarkFlagRequired("identity-source-arn")
	})
	mpaCmd.AddCommand(mpa_getIdentitySourceCmd)
}
