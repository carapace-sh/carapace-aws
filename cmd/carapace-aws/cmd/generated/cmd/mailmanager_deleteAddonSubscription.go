package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mailmanager_deleteAddonSubscriptionCmd = &cobra.Command{
	Use:   "delete-addon-subscription",
	Short: "Deletes an Add On subscription.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mailmanager_deleteAddonSubscriptionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mailmanager_deleteAddonSubscriptionCmd).Standalone()

		mailmanager_deleteAddonSubscriptionCmd.Flags().String("addon-subscription-id", "", "The Add On subscription ID to delete.")
		mailmanager_deleteAddonSubscriptionCmd.MarkFlagRequired("addon-subscription-id")
	})
	mailmanagerCmd.AddCommand(mailmanager_deleteAddonSubscriptionCmd)
}
