package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssmSap_startApplicationCmd = &cobra.Command{
	Use:   "start-application",
	Short: "Request is an operation which starts an application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssmSap_startApplicationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssmSap_startApplicationCmd).Standalone()

		ssmSap_startApplicationCmd.Flags().String("application-id", "", "The ID of the application.")
		ssmSap_startApplicationCmd.MarkFlagRequired("application-id")
	})
	ssmSapCmd.AddCommand(ssmSap_startApplicationCmd)
}
