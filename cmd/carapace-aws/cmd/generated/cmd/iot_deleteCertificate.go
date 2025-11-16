package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_deleteCertificateCmd = &cobra.Command{
	Use:   "delete-certificate",
	Short: "Deletes the specified certificate.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_deleteCertificateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_deleteCertificateCmd).Standalone()

		iot_deleteCertificateCmd.Flags().String("certificate-id", "", "The ID of the certificate.")
		iot_deleteCertificateCmd.Flags().String("force-delete", "", "Forces the deletion of a certificate if it is inactive and is not attached to an IoT thing.")
		iot_deleteCertificateCmd.MarkFlagRequired("certificate-id")
	})
	iotCmd.AddCommand(iot_deleteCertificateCmd)
}
