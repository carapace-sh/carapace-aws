package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_deleteCacertificateCmd = &cobra.Command{
	Use:   "delete-cacertificate",
	Short: "Deletes a registered CA certificate.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_deleteCacertificateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_deleteCacertificateCmd).Standalone()

		iot_deleteCacertificateCmd.Flags().String("certificate-id", "", "The ID of the certificate to delete.")
		iot_deleteCacertificateCmd.MarkFlagRequired("certificate-id")
	})
	iotCmd.AddCommand(iot_deleteCacertificateCmd)
}
