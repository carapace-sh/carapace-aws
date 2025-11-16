package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ds_verifyTrustCmd = &cobra.Command{
	Use:   "verify-trust",
	Short: "Directory Service for Microsoft Active Directory allows you to configure and verify trust relationships.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ds_verifyTrustCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ds_verifyTrustCmd).Standalone()

		ds_verifyTrustCmd.Flags().String("trust-id", "", "The unique Trust ID of the trust relationship to verify.")
		ds_verifyTrustCmd.MarkFlagRequired("trust-id")
	})
	dsCmd.AddCommand(ds_verifyTrustCmd)
}
