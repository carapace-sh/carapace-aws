package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rolesanywhere_getCrlCmd = &cobra.Command{
	Use:   "get-crl",
	Short: "Gets a certificate revocation list (CRL).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rolesanywhere_getCrlCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rolesanywhere_getCrlCmd).Standalone()

		rolesanywhere_getCrlCmd.Flags().String("crl-id", "", "The unique identifier of the certificate revocation list (CRL).")
		rolesanywhere_getCrlCmd.MarkFlagRequired("crl-id")
	})
	rolesanywhereCmd.AddCommand(rolesanywhere_getCrlCmd)
}
