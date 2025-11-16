package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpoint_listTemplateVersionsCmd = &cobra.Command{
	Use:   "list-template-versions",
	Short: "Retrieves information about all the versions of a specific message template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpoint_listTemplateVersionsCmd).Standalone()

	pinpoint_listTemplateVersionsCmd.Flags().String("next-token", "", "The string that specifies which page of results to return in a paginated response.")
	pinpoint_listTemplateVersionsCmd.Flags().String("page-size", "", "The maximum number of items to include in each page of a paginated response.")
	pinpoint_listTemplateVersionsCmd.Flags().String("template-name", "", "The name of the message template.")
	pinpoint_listTemplateVersionsCmd.Flags().String("template-type", "", "The type of channel that the message template is designed for.")
	pinpoint_listTemplateVersionsCmd.MarkFlagRequired("template-name")
	pinpoint_listTemplateVersionsCmd.MarkFlagRequired("template-type")
	pinpointCmd.AddCommand(pinpoint_listTemplateVersionsCmd)
}
