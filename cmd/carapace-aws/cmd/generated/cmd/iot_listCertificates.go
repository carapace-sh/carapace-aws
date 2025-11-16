package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_listCertificatesCmd = &cobra.Command{
	Use:   "list-certificates",
	Short: "Lists the certificates registered in your Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_listCertificatesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_listCertificatesCmd).Standalone()

		iot_listCertificatesCmd.Flags().String("ascending-order", "", "Specifies the order for results.")
		iot_listCertificatesCmd.Flags().String("marker", "", "The marker for the next set of results.")
		iot_listCertificatesCmd.Flags().String("page-size", "", "The result page size.")
	})
	iotCmd.AddCommand(iot_listCertificatesCmd)
}
