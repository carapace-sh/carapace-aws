package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_createCertificateFromCsrCmd = &cobra.Command{
	Use:   "create-certificate-from-csr",
	Short: "Creates an X.509 certificate using the specified certificate signing request.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_createCertificateFromCsrCmd).Standalone()

	iot_createCertificateFromCsrCmd.Flags().String("certificate-signing-request", "", "The certificate signing request (CSR).")
	iot_createCertificateFromCsrCmd.Flags().String("set-as-active", "", "Specifies whether the certificate is active.")
	iot_createCertificateFromCsrCmd.MarkFlagRequired("certificate-signing-request")
	iotCmd.AddCommand(iot_createCertificateFromCsrCmd)
}
