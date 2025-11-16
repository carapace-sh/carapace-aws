package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssmSap_deregisterApplicationCmd = &cobra.Command{
	Use:   "deregister-application",
	Short: "Deregister an SAP application with AWS Systems Manager for SAP.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssmSap_deregisterApplicationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssmSap_deregisterApplicationCmd).Standalone()

		ssmSap_deregisterApplicationCmd.Flags().String("application-id", "", "The ID of the application.")
		ssmSap_deregisterApplicationCmd.MarkFlagRequired("application-id")
	})
	ssmSapCmd.AddCommand(ssmSap_deregisterApplicationCmd)
}
