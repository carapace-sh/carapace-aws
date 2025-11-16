package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backupGateway_listGatewaysCmd = &cobra.Command{
	Use:   "list-gateways",
	Short: "Lists backup gateways owned by an Amazon Web Services account in an Amazon Web Services Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backupGateway_listGatewaysCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(backupGateway_listGatewaysCmd).Standalone()

		backupGateway_listGatewaysCmd.Flags().String("max-results", "", "The maximum number of gateways to list.")
		backupGateway_listGatewaysCmd.Flags().String("next-token", "", "The next item following a partial list of returned resources.")
	})
	backupGatewayCmd.AddCommand(backupGateway_listGatewaysCmd)
}
