package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var transfer_listAccessesCmd = &cobra.Command{
	Use:   "list-accesses",
	Short: "Lists the details for all the accesses you have on your server.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(transfer_listAccessesCmd).Standalone()

	transfer_listAccessesCmd.Flags().String("max-results", "", "The maximum number of items to return.")
	transfer_listAccessesCmd.Flags().String("next-token", "", "When you can get additional results from the `ListAccesses` call, a `NextToken` parameter is returned in the output.")
	transfer_listAccessesCmd.Flags().String("server-id", "", "A system-assigned unique identifier for a server that has users assigned to it.")
	transfer_listAccessesCmd.MarkFlagRequired("server-id")
	transferCmd.AddCommand(transfer_listAccessesCmd)
}
