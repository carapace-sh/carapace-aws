package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_updateCacertificateCmd = &cobra.Command{
	Use:   "update-cacertificate",
	Short: "Updates a registered CA certificate.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_updateCacertificateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_updateCacertificateCmd).Standalone()

		iot_updateCacertificateCmd.Flags().String("certificate-id", "", "The CA certificate identifier.")
		iot_updateCacertificateCmd.Flags().String("new-auto-registration-status", "", "The new value for the auto registration status.")
		iot_updateCacertificateCmd.Flags().String("new-status", "", "The updated status of the CA certificate.")
		iot_updateCacertificateCmd.Flags().String("registration-config", "", "Information about the registration configuration.")
		iot_updateCacertificateCmd.Flags().String("remove-auto-registration", "", "If true, removes auto registration.")
		iot_updateCacertificateCmd.MarkFlagRequired("certificate-id")
	})
	iotCmd.AddCommand(iot_updateCacertificateCmd)
}
