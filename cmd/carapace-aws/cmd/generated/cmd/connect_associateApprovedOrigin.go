package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_associateApprovedOriginCmd = &cobra.Command{
	Use:   "associate-approved-origin",
	Short: "This API is in preview release for Amazon Connect and is subject to change.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_associateApprovedOriginCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_associateApprovedOriginCmd).Standalone()

		connect_associateApprovedOriginCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		connect_associateApprovedOriginCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_associateApprovedOriginCmd.Flags().String("origin", "", "The domain to add to your allow list.")
		connect_associateApprovedOriginCmd.MarkFlagRequired("instance-id")
		connect_associateApprovedOriginCmd.MarkFlagRequired("origin")
	})
	connectCmd.AddCommand(connect_associateApprovedOriginCmd)
}
