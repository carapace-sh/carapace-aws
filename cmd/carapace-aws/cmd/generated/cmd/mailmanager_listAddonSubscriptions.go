package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mailmanager_listAddonSubscriptionsCmd = &cobra.Command{
	Use:   "list-addon-subscriptions",
	Short: "Lists all Add On subscriptions in your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mailmanager_listAddonSubscriptionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mailmanager_listAddonSubscriptionsCmd).Standalone()

		mailmanager_listAddonSubscriptionsCmd.Flags().String("next-token", "", "If you received a pagination token from a previous call to this API, you can provide it here to continue paginating through the next page of results.")
		mailmanager_listAddonSubscriptionsCmd.Flags().String("page-size", "", "The maximum number of ingress endpoint resources that are returned per call.")
	})
	mailmanagerCmd.AddCommand(mailmanager_listAddonSubscriptionsCmd)
}
