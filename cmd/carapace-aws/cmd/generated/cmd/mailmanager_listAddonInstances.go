package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mailmanager_listAddonInstancesCmd = &cobra.Command{
	Use:   "list-addon-instances",
	Short: "Lists all Add On instances in your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mailmanager_listAddonInstancesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mailmanager_listAddonInstancesCmd).Standalone()

		mailmanager_listAddonInstancesCmd.Flags().String("next-token", "", "If you received a pagination token from a previous call to this API, you can provide it here to continue paginating through the next page of results.")
		mailmanager_listAddonInstancesCmd.Flags().String("page-size", "", "The maximum number of ingress endpoint resources that are returned per call.")
	})
	mailmanagerCmd.AddCommand(mailmanager_listAddonInstancesCmd)
}
