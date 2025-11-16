package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointEmail_listEmailIdentitiesCmd = &cobra.Command{
	Use:   "list-email-identities",
	Short: "Returns a list of all of the email identities that are associated with your Amazon Pinpoint account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointEmail_listEmailIdentitiesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pinpointEmail_listEmailIdentitiesCmd).Standalone()

		pinpointEmail_listEmailIdentitiesCmd.Flags().String("next-token", "", "A token returned from a previous call to `ListEmailIdentities` to indicate the position in the list of identities.")
		pinpointEmail_listEmailIdentitiesCmd.Flags().String("page-size", "", "The number of results to show in a single call to `ListEmailIdentities`.")
	})
	pinpointEmailCmd.AddCommand(pinpointEmail_listEmailIdentitiesCmd)
}
