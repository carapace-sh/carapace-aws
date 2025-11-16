package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpoint_createInAppTemplateCmd = &cobra.Command{
	Use:   "create-in-app-template",
	Short: "Creates a new message template for messages using the in-app message channel.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpoint_createInAppTemplateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pinpoint_createInAppTemplateCmd).Standalone()

		pinpoint_createInAppTemplateCmd.Flags().String("in-app-template-request", "", "")
		pinpoint_createInAppTemplateCmd.Flags().String("template-name", "", "The name of the message template.")
		pinpoint_createInAppTemplateCmd.MarkFlagRequired("in-app-template-request")
		pinpoint_createInAppTemplateCmd.MarkFlagRequired("template-name")
	})
	pinpointCmd.AddCommand(pinpoint_createInAppTemplateCmd)
}
