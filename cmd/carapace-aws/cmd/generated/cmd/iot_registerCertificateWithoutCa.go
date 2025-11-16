package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_registerCertificateWithoutCaCmd = &cobra.Command{
	Use:   "register-certificate-without-ca",
	Short: "Register a certificate that does not have a certificate authority (CA).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_registerCertificateWithoutCaCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_registerCertificateWithoutCaCmd).Standalone()

		iot_registerCertificateWithoutCaCmd.Flags().String("certificate-pem", "", "The certificate data, in PEM format.")
		iot_registerCertificateWithoutCaCmd.Flags().String("status", "", "The status of the register certificate request.")
		iot_registerCertificateWithoutCaCmd.MarkFlagRequired("certificate-pem")
	})
	iotCmd.AddCommand(iot_registerCertificateWithoutCaCmd)
}
