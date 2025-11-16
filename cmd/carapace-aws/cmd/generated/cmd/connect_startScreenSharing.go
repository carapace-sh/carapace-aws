package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_startScreenSharingCmd = &cobra.Command{
	Use:   "start-screen-sharing",
	Short: "Starts screen sharing for a contact.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_startScreenSharingCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_startScreenSharingCmd).Standalone()

		connect_startScreenSharingCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		connect_startScreenSharingCmd.Flags().String("contact-id", "", "The identifier of the contact in this instance of Amazon Connect.")
		connect_startScreenSharingCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_startScreenSharingCmd.MarkFlagRequired("contact-id")
		connect_startScreenSharingCmd.MarkFlagRequired("instance-id")
	})
	connectCmd.AddCommand(connect_startScreenSharingCmd)
}
