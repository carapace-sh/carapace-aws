package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpoint_deleteInAppTemplateCmd = &cobra.Command{
	Use:   "delete-in-app-template",
	Short: "Deletes a message template for messages sent using the in-app message channel.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpoint_deleteInAppTemplateCmd).Standalone()

	pinpoint_deleteInAppTemplateCmd.Flags().String("template-name", "", "The name of the message template.")
	pinpoint_deleteInAppTemplateCmd.Flags().String("version", "", "The unique identifier for the version of the message template to update, retrieve information about, or delete.")
	pinpoint_deleteInAppTemplateCmd.MarkFlagRequired("template-name")
	pinpointCmd.AddCommand(pinpoint_deleteInAppTemplateCmd)
}
