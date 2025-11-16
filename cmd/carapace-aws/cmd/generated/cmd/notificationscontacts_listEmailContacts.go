package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var notificationscontacts_listEmailContactsCmd = &cobra.Command{
	Use:   "list-email-contacts",
	Short: "Lists all email contacts created under the Account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(notificationscontacts_listEmailContactsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(notificationscontacts_listEmailContactsCmd).Standalone()

		notificationscontacts_listEmailContactsCmd.Flags().String("max-results", "", "The maximum number of results to include in the response.")
		notificationscontacts_listEmailContactsCmd.Flags().String("next-token", "", "An optional token returned from a prior request.")
	})
	notificationscontactsCmd.AddCommand(notificationscontacts_listEmailContactsCmd)
}
