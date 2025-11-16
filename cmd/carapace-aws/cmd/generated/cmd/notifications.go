package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var notificationsCmd = &cobra.Command{
	Use:   "notifications",
	Short: "The *User Notifications API Reference* provides descriptions, API request parameters, and the JSON response for each of the User Notifications API actions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(notificationsCmd).Standalone()

	rootCmd.AddCommand(notificationsCmd)
}
