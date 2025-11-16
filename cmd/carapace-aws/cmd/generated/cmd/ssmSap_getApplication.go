package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssmSap_getApplicationCmd = &cobra.Command{
	Use:   "get-application",
	Short: "Gets an application registered with AWS Systems Manager for SAP.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssmSap_getApplicationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssmSap_getApplicationCmd).Standalone()

		ssmSap_getApplicationCmd.Flags().String("app-registry-arn", "", "The Amazon Resource Name (ARN) of the application registry.")
		ssmSap_getApplicationCmd.Flags().String("application-arn", "", "The Amazon Resource Name (ARN) of the application.")
		ssmSap_getApplicationCmd.Flags().String("application-id", "", "The ID of the application.")
	})
	ssmSapCmd.AddCommand(ssmSap_getApplicationCmd)
}
