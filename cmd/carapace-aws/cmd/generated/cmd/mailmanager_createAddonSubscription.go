package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mailmanager_createAddonSubscriptionCmd = &cobra.Command{
	Use:   "create-addon-subscription",
	Short: "Creates a subscription for an Add On representing the acceptance of its terms of use and additional pricing.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mailmanager_createAddonSubscriptionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mailmanager_createAddonSubscriptionCmd).Standalone()

		mailmanager_createAddonSubscriptionCmd.Flags().String("addon-name", "", "The name of the Add On to subscribe to.")
		mailmanager_createAddonSubscriptionCmd.Flags().String("client-token", "", "A unique token that Amazon SES uses to recognize subsequent retries of the same request.")
		mailmanager_createAddonSubscriptionCmd.Flags().String("tags", "", "The tags used to organize, track, or control access for the resource.")
		mailmanager_createAddonSubscriptionCmd.MarkFlagRequired("addon-name")
	})
	mailmanagerCmd.AddCommand(mailmanager_createAddonSubscriptionCmd)
}
