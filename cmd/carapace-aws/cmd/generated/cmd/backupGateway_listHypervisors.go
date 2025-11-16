package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backupGateway_listHypervisorsCmd = &cobra.Command{
	Use:   "list-hypervisors",
	Short: "Lists your hypervisors.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backupGateway_listHypervisorsCmd).Standalone()

	backupGateway_listHypervisorsCmd.Flags().String("max-results", "", "The maximum number of hypervisors to list.")
	backupGateway_listHypervisorsCmd.Flags().String("next-token", "", "The next item following a partial list of returned resources.")
	backupGatewayCmd.AddCommand(backupGateway_listHypervisorsCmd)
}
