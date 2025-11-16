package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_updateCertificateCmd = &cobra.Command{
	Use:   "update-certificate",
	Short: "Updates the status of the specified certificate.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_updateCertificateCmd).Standalone()

	iot_updateCertificateCmd.Flags().String("certificate-id", "", "The ID of the certificate.")
	iot_updateCertificateCmd.Flags().String("new-status", "", "The new status.")
	iot_updateCertificateCmd.MarkFlagRequired("certificate-id")
	iot_updateCertificateCmd.MarkFlagRequired("new-status")
	iotCmd.AddCommand(iot_updateCertificateCmd)
}
