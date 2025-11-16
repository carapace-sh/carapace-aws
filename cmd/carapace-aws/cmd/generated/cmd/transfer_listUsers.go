package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var transfer_listUsersCmd = &cobra.Command{
	Use:   "list-users",
	Short: "Lists the users for a file transfer protocol-enabled server that you specify by passing the `ServerId` parameter.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(transfer_listUsersCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(transfer_listUsersCmd).Standalone()

		transfer_listUsersCmd.Flags().String("max-results", "", "Specifies the number of users to return as a response to the `ListUsers` request.")
		transfer_listUsersCmd.Flags().String("next-token", "", "If there are additional results from the `ListUsers` call, a `NextToken` parameter is returned in the output.")
		transfer_listUsersCmd.Flags().String("server-id", "", "A system-assigned unique identifier for a server that has users assigned to it.")
		transfer_listUsersCmd.MarkFlagRequired("server-id")
	})
	transferCmd.AddCommand(transfer_listUsersCmd)
}
