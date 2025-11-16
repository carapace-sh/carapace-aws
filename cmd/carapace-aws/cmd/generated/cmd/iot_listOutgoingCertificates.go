package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_listOutgoingCertificatesCmd = &cobra.Command{
	Use:   "list-outgoing-certificates",
	Short: "Lists certificates that are being transferred but not yet accepted.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_listOutgoingCertificatesCmd).Standalone()

	iot_listOutgoingCertificatesCmd.Flags().String("ascending-order", "", "Specifies the order for results.")
	iot_listOutgoingCertificatesCmd.Flags().String("marker", "", "The marker for the next set of results.")
	iot_listOutgoingCertificatesCmd.Flags().String("page-size", "", "The result page size.")
	iotCmd.AddCommand(iot_listOutgoingCertificatesCmd)
}
