package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ses_getIdentityNotificationAttributesCmd = &cobra.Command{
	Use:   "get-identity-notification-attributes",
	Short: "Given a list of verified identities (email addresses and/or domains), returns a structure describing identity notification attributes.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ses_getIdentityNotificationAttributesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ses_getIdentityNotificationAttributesCmd).Standalone()

		ses_getIdentityNotificationAttributesCmd.Flags().String("identities", "", "A list of one or more identities.")
		ses_getIdentityNotificationAttributesCmd.MarkFlagRequired("identities")
	})
	sesCmd.AddCommand(ses_getIdentityNotificationAttributesCmd)
}
