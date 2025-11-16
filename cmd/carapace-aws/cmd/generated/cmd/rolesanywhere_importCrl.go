package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rolesanywhere_importCrlCmd = &cobra.Command{
	Use:   "import-crl",
	Short: "Imports the certificate revocation list (CRL).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rolesanywhere_importCrlCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rolesanywhere_importCrlCmd).Standalone()

		rolesanywhere_importCrlCmd.Flags().String("crl-data", "", "The x509 v3 specified certificate revocation list (CRL).")
		rolesanywhere_importCrlCmd.Flags().Bool("enabled", false, "Specifies whether the certificate revocation list (CRL) is enabled.")
		rolesanywhere_importCrlCmd.Flags().String("name", "", "The name of the certificate revocation list (CRL).")
		rolesanywhere_importCrlCmd.Flags().Bool("no-enabled", false, "Specifies whether the certificate revocation list (CRL) is enabled.")
		rolesanywhere_importCrlCmd.Flags().String("tags", "", "A list of tags to attach to the certificate revocation list (CRL).")
		rolesanywhere_importCrlCmd.Flags().String("trust-anchor-arn", "", "The ARN of the TrustAnchor the certificate revocation list (CRL) will provide revocation for.")
		rolesanywhere_importCrlCmd.MarkFlagRequired("crl-data")
		rolesanywhere_importCrlCmd.MarkFlagRequired("name")
		rolesanywhere_importCrlCmd.Flag("no-enabled").Hidden = true
		rolesanywhere_importCrlCmd.MarkFlagRequired("trust-anchor-arn")
	})
	rolesanywhereCmd.AddCommand(rolesanywhere_importCrlCmd)
}
