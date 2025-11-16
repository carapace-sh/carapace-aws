package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_associateSecurityKeyCmd = &cobra.Command{
	Use:   "associate-security-key",
	Short: "This API is in preview release for Amazon Connect and is subject to change.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_associateSecurityKeyCmd).Standalone()

	connect_associateSecurityKeyCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	connect_associateSecurityKeyCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_associateSecurityKeyCmd.Flags().String("key", "", "A valid security key in PEM format as a String.")
	connect_associateSecurityKeyCmd.MarkFlagRequired("instance-id")
	connect_associateSecurityKeyCmd.MarkFlagRequired("key")
	connectCmd.AddCommand(connect_associateSecurityKeyCmd)
}
