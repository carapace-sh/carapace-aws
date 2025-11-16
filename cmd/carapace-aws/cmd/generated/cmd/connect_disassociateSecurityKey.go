package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_disassociateSecurityKeyCmd = &cobra.Command{
	Use:   "disassociate-security-key",
	Short: "This API is in preview release for Amazon Connect and is subject to change.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_disassociateSecurityKeyCmd).Standalone()

	connect_disassociateSecurityKeyCmd.Flags().String("association-id", "", "The existing association identifier that uniquely identifies the resource type and storage config for the given instance ID.")
	connect_disassociateSecurityKeyCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	connect_disassociateSecurityKeyCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_disassociateSecurityKeyCmd.MarkFlagRequired("association-id")
	connect_disassociateSecurityKeyCmd.MarkFlagRequired("instance-id")
	connectCmd.AddCommand(connect_disassociateSecurityKeyCmd)
}
