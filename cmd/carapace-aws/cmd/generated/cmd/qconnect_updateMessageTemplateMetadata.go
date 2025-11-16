package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qconnect_updateMessageTemplateMetadataCmd = &cobra.Command{
	Use:   "update-message-template-metadata",
	Short: "Updates the Amazon Q in Connect message template metadata.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qconnect_updateMessageTemplateMetadataCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(qconnect_updateMessageTemplateMetadataCmd).Standalone()

		qconnect_updateMessageTemplateMetadataCmd.Flags().String("description", "", "The description of the message template.")
		qconnect_updateMessageTemplateMetadataCmd.Flags().String("grouping-configuration", "", "")
		qconnect_updateMessageTemplateMetadataCmd.Flags().String("knowledge-base-id", "", "The identifier of the knowledge base.")
		qconnect_updateMessageTemplateMetadataCmd.Flags().String("message-template-id", "", "The identifier of the message template.")
		qconnect_updateMessageTemplateMetadataCmd.Flags().String("name", "", "The name of the message template.")
		qconnect_updateMessageTemplateMetadataCmd.MarkFlagRequired("knowledge-base-id")
		qconnect_updateMessageTemplateMetadataCmd.MarkFlagRequired("message-template-id")
	})
	qconnectCmd.AddCommand(qconnect_updateMessageTemplateMetadataCmd)
}
