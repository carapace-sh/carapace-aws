package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_listCacertificatesCmd = &cobra.Command{
	Use:   "list-cacertificates",
	Short: "Lists the CA certificates registered for your Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_listCacertificatesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_listCacertificatesCmd).Standalone()

		iot_listCacertificatesCmd.Flags().String("ascending-order", "", "Determines the order of the results.")
		iot_listCacertificatesCmd.Flags().String("marker", "", "The marker for the next set of results.")
		iot_listCacertificatesCmd.Flags().String("page-size", "", "The result page size.")
		iot_listCacertificatesCmd.Flags().String("template-name", "", "The name of the provisioning template.")
	})
	iotCmd.AddCommand(iot_listCacertificatesCmd)
}
