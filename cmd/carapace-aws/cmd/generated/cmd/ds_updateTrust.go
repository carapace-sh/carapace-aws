package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ds_updateTrustCmd = &cobra.Command{
	Use:   "update-trust",
	Short: "Updates the trust that has been set up between your Managed Microsoft AD directory and an self-managed Active Directory.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ds_updateTrustCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ds_updateTrustCmd).Standalone()

		ds_updateTrustCmd.Flags().String("selective-auth", "", "Updates selective authentication for the trust.")
		ds_updateTrustCmd.Flags().String("trust-id", "", "Identifier of the trust relationship.")
		ds_updateTrustCmd.MarkFlagRequired("trust-id")
	})
	dsCmd.AddCommand(ds_updateTrustCmd)
}
