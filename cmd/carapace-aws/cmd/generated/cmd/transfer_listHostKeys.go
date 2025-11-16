package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var transfer_listHostKeysCmd = &cobra.Command{
	Use:   "list-host-keys",
	Short: "Returns a list of host keys for the server that's specified by the `ServerId` parameter.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(transfer_listHostKeysCmd).Standalone()

	transfer_listHostKeysCmd.Flags().String("max-results", "", "The maximum number of items to return.")
	transfer_listHostKeysCmd.Flags().String("next-token", "", "When there are additional results that were not returned, a `NextToken` parameter is returned.")
	transfer_listHostKeysCmd.Flags().String("server-id", "", "The identifier of the server that contains the host keys that you want to view.")
	transfer_listHostKeysCmd.MarkFlagRequired("server-id")
	transferCmd.AddCommand(transfer_listHostKeysCmd)
}
