package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_disassociateApprovedOriginCmd = &cobra.Command{
	Use:   "disassociate-approved-origin",
	Short: "This API is in preview release for Amazon Connect and is subject to change.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_disassociateApprovedOriginCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_disassociateApprovedOriginCmd).Standalone()

		connect_disassociateApprovedOriginCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		connect_disassociateApprovedOriginCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_disassociateApprovedOriginCmd.Flags().String("origin", "", "The domain URL of the integrated application.")
		connect_disassociateApprovedOriginCmd.MarkFlagRequired("instance-id")
		connect_disassociateApprovedOriginCmd.MarkFlagRequired("origin")
	})
	connectCmd.AddCommand(connect_disassociateApprovedOriginCmd)
}
