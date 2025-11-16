package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpoint_listTemplatesCmd = &cobra.Command{
	Use:   "list-templates",
	Short: "Retrieves information about all the message templates that are associated with your Amazon Pinpoint account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpoint_listTemplatesCmd).Standalone()

	pinpoint_listTemplatesCmd.Flags().String("next-token", "", "The string that specifies which page of results to return in a paginated response.")
	pinpoint_listTemplatesCmd.Flags().String("page-size", "", "The maximum number of items to include in each page of a paginated response.")
	pinpoint_listTemplatesCmd.Flags().String("prefix", "", "The substring to match in the names of the message templates to include in the results.")
	pinpoint_listTemplatesCmd.Flags().String("template-type", "", "The type of message template to include in the results.")
	pinpointCmd.AddCommand(pinpoint_listTemplatesCmd)
}
