package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mailmanager_listRelaysCmd = &cobra.Command{
	Use:   "list-relays",
	Short: "Lists all the existing relay resources.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mailmanager_listRelaysCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mailmanager_listRelaysCmd).Standalone()

		mailmanager_listRelaysCmd.Flags().String("next-token", "", "If you received a pagination token from a previous call to this API, you can provide it here to continue paginating through the next page of results.")
		mailmanager_listRelaysCmd.Flags().String("page-size", "", "The number of relays to be returned in one request.")
	})
	mailmanagerCmd.AddCommand(mailmanager_listRelaysCmd)
}
