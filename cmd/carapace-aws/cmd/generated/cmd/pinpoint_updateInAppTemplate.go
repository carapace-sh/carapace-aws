package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpoint_updateInAppTemplateCmd = &cobra.Command{
	Use:   "update-in-app-template",
	Short: "Updates an existing message template for messages sent through the in-app message channel.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpoint_updateInAppTemplateCmd).Standalone()

	pinpoint_updateInAppTemplateCmd.Flags().String("create-new-version", "", "Specifies whether to save the updates as a new version of the message template.")
	pinpoint_updateInAppTemplateCmd.Flags().String("in-app-template-request", "", "")
	pinpoint_updateInAppTemplateCmd.Flags().String("template-name", "", "The name of the message template.")
	pinpoint_updateInAppTemplateCmd.Flags().String("version", "", "The unique identifier for the version of the message template to update, retrieve information about, or delete.")
	pinpoint_updateInAppTemplateCmd.MarkFlagRequired("in-app-template-request")
	pinpoint_updateInAppTemplateCmd.MarkFlagRequired("template-name")
	pinpointCmd.AddCommand(pinpoint_updateInAppTemplateCmd)
}
