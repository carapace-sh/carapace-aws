package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var transfer_listServersCmd = &cobra.Command{
	Use:   "list-servers",
	Short: "Lists the file transfer protocol-enabled servers that are associated with your Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(transfer_listServersCmd).Standalone()

	transfer_listServersCmd.Flags().String("max-results", "", "Specifies the number of servers to return as a response to the `ListServers` query.")
	transfer_listServersCmd.Flags().String("next-token", "", "When additional results are obtained from the `ListServers` command, a `NextToken` parameter is returned in the output.")
	transferCmd.AddCommand(transfer_listServersCmd)
}
