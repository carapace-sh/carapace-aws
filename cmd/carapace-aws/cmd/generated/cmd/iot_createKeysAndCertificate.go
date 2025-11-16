package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_createKeysAndCertificateCmd = &cobra.Command{
	Use:   "create-keys-and-certificate",
	Short: "Creates a 2048-bit RSA key pair and issues an X.509 certificate using the issued public key.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_createKeysAndCertificateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_createKeysAndCertificateCmd).Standalone()

		iot_createKeysAndCertificateCmd.Flags().String("set-as-active", "", "Specifies whether the certificate is active.")
	})
	iotCmd.AddCommand(iot_createKeysAndCertificateCmd)
}
