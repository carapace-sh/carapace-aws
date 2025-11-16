package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_describeCacertificateCmd = &cobra.Command{
	Use:   "describe-cacertificate",
	Short: "Describes a registered CA certificate.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_describeCacertificateCmd).Standalone()

	iot_describeCacertificateCmd.Flags().String("certificate-id", "", "The CA certificate identifier.")
	iot_describeCacertificateCmd.MarkFlagRequired("certificate-id")
	iotCmd.AddCommand(iot_describeCacertificateCmd)
}
