package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qconnect_createAssistantCmd = &cobra.Command{
	Use:   "create-assistant",
	Short: "Creates an Amazon Q in Connect assistant.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qconnect_createAssistantCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(qconnect_createAssistantCmd).Standalone()

		qconnect_createAssistantCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		qconnect_createAssistantCmd.Flags().String("description", "", "The description of the assistant.")
		qconnect_createAssistantCmd.Flags().String("name", "", "The name of the assistant.")
		qconnect_createAssistantCmd.Flags().String("server-side-encryption-configuration", "", "The configuration information for the customer managed key used for encryption.")
		qconnect_createAssistantCmd.Flags().String("tags", "", "The tags used to organize, track, or control access for this resource.")
		qconnect_createAssistantCmd.Flags().String("type", "", "The type of assistant.")
		qconnect_createAssistantCmd.MarkFlagRequired("name")
		qconnect_createAssistantCmd.MarkFlagRequired("type")
	})
	qconnectCmd.AddCommand(qconnect_createAssistantCmd)
}
