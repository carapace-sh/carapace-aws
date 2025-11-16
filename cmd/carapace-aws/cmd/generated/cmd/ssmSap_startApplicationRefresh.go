package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssmSap_startApplicationRefreshCmd = &cobra.Command{
	Use:   "start-application-refresh",
	Short: "Refreshes a registered application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssmSap_startApplicationRefreshCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssmSap_startApplicationRefreshCmd).Standalone()

		ssmSap_startApplicationRefreshCmd.Flags().String("application-id", "", "The ID of the application.")
		ssmSap_startApplicationRefreshCmd.MarkFlagRequired("application-id")
	})
	ssmSapCmd.AddCommand(ssmSap_startApplicationRefreshCmd)
}
