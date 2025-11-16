package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rolesanywhere_disableCrlCmd = &cobra.Command{
	Use:   "disable-crl",
	Short: "Disables a certificate revocation list (CRL).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rolesanywhere_disableCrlCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rolesanywhere_disableCrlCmd).Standalone()

		rolesanywhere_disableCrlCmd.Flags().String("crl-id", "", "The unique identifier of the certificate revocation list (CRL).")
		rolesanywhere_disableCrlCmd.MarkFlagRequired("crl-id")
	})
	rolesanywhereCmd.AddCommand(rolesanywhere_disableCrlCmd)
}
