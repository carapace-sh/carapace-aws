package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var notificationscontacts_activateEmailContactCmd = &cobra.Command{
	Use:   "activate-email-contact",
	Short: "Activates an email contact using an activation code.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(notificationscontacts_activateEmailContactCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(notificationscontacts_activateEmailContactCmd).Standalone()

		notificationscontacts_activateEmailContactCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) of the resource.")
		notificationscontacts_activateEmailContactCmd.Flags().String("code", "", "The activation code for this email contact.")
		notificationscontacts_activateEmailContactCmd.MarkFlagRequired("arn")
		notificationscontacts_activateEmailContactCmd.MarkFlagRequired("code")
	})
	notificationscontactsCmd.AddCommand(notificationscontacts_activateEmailContactCmd)
}
