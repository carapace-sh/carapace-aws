package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var notificationscontactsCmd = &cobra.Command{
	Use:   "notificationscontacts",
	Short: "AWS User Notifications Contacts is a service that allows you to create and manage email contacts for AWS User Notifications.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(notificationscontactsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(notificationscontactsCmd).Standalone()

	})
	rootCmd.AddCommand(notificationscontactsCmd)
}
