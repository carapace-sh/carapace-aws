package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_registerCacertificateCmd = &cobra.Command{
	Use:   "register-cacertificate",
	Short: "Registers a CA certificate with Amazon Web Services IoT Core.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_registerCacertificateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_registerCacertificateCmd).Standalone()

		iot_registerCacertificateCmd.Flags().String("allow-auto-registration", "", "Allows this CA certificate to be used for auto registration of device certificates.")
		iot_registerCacertificateCmd.Flags().String("ca-certificate", "", "The CA certificate.")
		iot_registerCacertificateCmd.Flags().String("certificate-mode", "", "Describes the certificate mode in which the Certificate Authority (CA) will be registered.")
		iot_registerCacertificateCmd.Flags().String("registration-config", "", "Information about the registration configuration.")
		iot_registerCacertificateCmd.Flags().String("set-as-active", "", "A boolean value that specifies if the CA certificate is set to active.")
		iot_registerCacertificateCmd.Flags().String("tags", "", "Metadata which can be used to manage the CA certificate.")
		iot_registerCacertificateCmd.Flags().String("verification-certificate", "", "The private key verification certificate.")
		iot_registerCacertificateCmd.MarkFlagRequired("ca-certificate")
	})
	iotCmd.AddCommand(iot_registerCacertificateCmd)
}
