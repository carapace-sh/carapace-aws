package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wisdom_createAssistantCmd = &cobra.Command{
	Use:   "create-assistant",
	Short: "Creates an Amazon Connect Wisdom assistant.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wisdom_createAssistantCmd).Standalone()

	wisdom_createAssistantCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	wisdom_createAssistantCmd.Flags().String("description", "", "The description of the assistant.")
	wisdom_createAssistantCmd.Flags().String("name", "", "The name of the assistant.")
	wisdom_createAssistantCmd.Flags().String("server-side-encryption-configuration", "", "The configuration information for the customer managed key used for encryption.")
	wisdom_createAssistantCmd.Flags().String("tags", "", "The tags used to organize, track, or control access for this resource.")
	wisdom_createAssistantCmd.Flags().String("type", "", "The type of assistant.")
	wisdom_createAssistantCmd.MarkFlagRequired("name")
	wisdom_createAssistantCmd.MarkFlagRequired("type")
	wisdomCmd.AddCommand(wisdom_createAssistantCmd)
}
