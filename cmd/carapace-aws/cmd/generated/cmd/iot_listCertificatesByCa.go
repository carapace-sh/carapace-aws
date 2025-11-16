package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_listCertificatesByCaCmd = &cobra.Command{
	Use:   "list-certificates-by-ca",
	Short: "List the device certificates signed by the specified CA certificate.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_listCertificatesByCaCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_listCertificatesByCaCmd).Standalone()

		iot_listCertificatesByCaCmd.Flags().String("ascending-order", "", "Specifies the order for results.")
		iot_listCertificatesByCaCmd.Flags().String("ca-certificate-id", "", "The ID of the CA certificate.")
		iot_listCertificatesByCaCmd.Flags().String("marker", "", "The marker for the next set of results.")
		iot_listCertificatesByCaCmd.Flags().String("page-size", "", "The result page size.")
		iot_listCertificatesByCaCmd.MarkFlagRequired("ca-certificate-id")
	})
	iotCmd.AddCommand(iot_listCertificatesByCaCmd)
}
