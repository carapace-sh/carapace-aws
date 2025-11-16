package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpoint_updateTemplateActiveVersionCmd = &cobra.Command{
	Use:   "update-template-active-version",
	Short: "Changes the status of a specific version of a message template to *active*.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpoint_updateTemplateActiveVersionCmd).Standalone()

	pinpoint_updateTemplateActiveVersionCmd.Flags().String("template-active-version-request", "", "")
	pinpoint_updateTemplateActiveVersionCmd.Flags().String("template-name", "", "The name of the message template.")
	pinpoint_updateTemplateActiveVersionCmd.Flags().String("template-type", "", "The type of channel that the message template is designed for.")
	pinpoint_updateTemplateActiveVersionCmd.MarkFlagRequired("template-active-version-request")
	pinpoint_updateTemplateActiveVersionCmd.MarkFlagRequired("template-name")
	pinpoint_updateTemplateActiveVersionCmd.MarkFlagRequired("template-type")
	pinpointCmd.AddCommand(pinpoint_updateTemplateActiveVersionCmd)
}
