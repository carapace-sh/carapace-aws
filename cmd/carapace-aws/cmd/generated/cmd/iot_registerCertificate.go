package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_registerCertificateCmd = &cobra.Command{
	Use:   "register-certificate",
	Short: "Registers a device certificate with IoT in the same [certificate mode](https://docs.aws.amazon.com/iot/latest/apireference/API_CertificateDescription.html#iot-Type-CertificateDescription-certificateMode) as the signing CA.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_registerCertificateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_registerCertificateCmd).Standalone()

		iot_registerCertificateCmd.Flags().String("ca-certificate-pem", "", "The CA certificate used to sign the device certificate being registered.")
		iot_registerCertificateCmd.Flags().String("certificate-pem", "", "The certificate data, in PEM format.")
		iot_registerCertificateCmd.Flags().String("set-as-active", "", "A boolean value that specifies if the certificate is set to active.")
		iot_registerCertificateCmd.Flags().String("status", "", "The status of the register certificate request.")
		iot_registerCertificateCmd.MarkFlagRequired("certificate-pem")
	})
	iotCmd.AddCommand(iot_registerCertificateCmd)
}
