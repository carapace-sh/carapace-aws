package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpoint_getInAppTemplateCmd = &cobra.Command{
	Use:   "get-in-app-template",
	Short: "Retrieves the content and settings of a message template for messages sent through the in-app channel.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpoint_getInAppTemplateCmd).Standalone()

	pinpoint_getInAppTemplateCmd.Flags().String("template-name", "", "The name of the message template.")
	pinpoint_getInAppTemplateCmd.Flags().String("version", "", "The unique identifier for the version of the message template to update, retrieve information about, or delete.")
	pinpoint_getInAppTemplateCmd.MarkFlagRequired("template-name")
	pinpointCmd.AddCommand(pinpoint_getInAppTemplateCmd)
}
