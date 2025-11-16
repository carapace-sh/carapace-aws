package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_createInstanceCmd = &cobra.Command{
	Use:   "create-instance",
	Short: "This API is in preview release for Amazon Connect and is subject to change.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_createInstanceCmd).Standalone()

	connect_createInstanceCmd.Flags().String("client-token", "", "The idempotency token.")
	connect_createInstanceCmd.Flags().String("directory-id", "", "The identifier for the directory.")
	connect_createInstanceCmd.Flags().String("identity-management-type", "", "The type of identity management for your Amazon Connect users.")
	connect_createInstanceCmd.Flags().String("inbound-calls-enabled", "", "Your contact center handles incoming contacts.")
	connect_createInstanceCmd.Flags().String("instance-alias", "", "The name for your instance.")
	connect_createInstanceCmd.Flags().String("outbound-calls-enabled", "", "Your contact center allows outbound calls.")
	connect_createInstanceCmd.Flags().String("tags", "", "The tags used to organize, track, or control access for this resource.")
	connect_createInstanceCmd.MarkFlagRequired("identity-management-type")
	connect_createInstanceCmd.MarkFlagRequired("inbound-calls-enabled")
	connect_createInstanceCmd.MarkFlagRequired("outbound-calls-enabled")
	connectCmd.AddCommand(connect_createInstanceCmd)
}
