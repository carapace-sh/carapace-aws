package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mailmanager_getAddonSubscriptionCmd = &cobra.Command{
	Use:   "get-addon-subscription",
	Short: "Gets detailed information about an Add On subscription.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mailmanager_getAddonSubscriptionCmd).Standalone()

	mailmanager_getAddonSubscriptionCmd.Flags().String("addon-subscription-id", "", "The Add On subscription ID to retrieve information for.")
	mailmanager_getAddonSubscriptionCmd.MarkFlagRequired("addon-subscription-id")
	mailmanagerCmd.AddCommand(mailmanager_getAddonSubscriptionCmd)
}
