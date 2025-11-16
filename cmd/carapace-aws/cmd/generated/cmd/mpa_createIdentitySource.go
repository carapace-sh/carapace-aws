package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mpa_createIdentitySourceCmd = &cobra.Command{
	Use:   "create-identity-source",
	Short: "Creates a new identity source.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mpa_createIdentitySourceCmd).Standalone()

	mpa_createIdentitySourceCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	mpa_createIdentitySourceCmd.Flags().String("identity-source-parameters", "", "A `IdentitySourceParameters` object.")
	mpa_createIdentitySourceCmd.Flags().String("tags", "", "Tag you want to attach to the identity source.")
	mpa_createIdentitySourceCmd.MarkFlagRequired("identity-source-parameters")
	mpaCmd.AddCommand(mpa_createIdentitySourceCmd)
}
