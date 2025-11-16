package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_listCertificateProvidersCmd = &cobra.Command{
	Use:   "list-certificate-providers",
	Short: "Lists all your certificate providers in your Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_listCertificateProvidersCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_listCertificateProvidersCmd).Standalone()

		iot_listCertificateProvidersCmd.Flags().String("ascending-order", "", "Returns the list of certificate providers in ascending alphabetical order.")
		iot_listCertificateProvidersCmd.Flags().String("next-token", "", "The token for the next set of results, or `null` if there are no more results.")
	})
	iotCmd.AddCommand(iot_listCertificateProvidersCmd)
}
