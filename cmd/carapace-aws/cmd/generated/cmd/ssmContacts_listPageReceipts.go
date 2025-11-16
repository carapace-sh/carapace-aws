package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssmContacts_listPageReceiptsCmd = &cobra.Command{
	Use:   "list-page-receipts",
	Short: "Lists all of the engagements to contact channels that have been acknowledged.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssmContacts_listPageReceiptsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssmContacts_listPageReceiptsCmd).Standalone()

		ssmContacts_listPageReceiptsCmd.Flags().String("max-results", "", "The maximum number of acknowledgements per page of results.")
		ssmContacts_listPageReceiptsCmd.Flags().String("next-token", "", "The pagination token to continue to the next page of results.")
		ssmContacts_listPageReceiptsCmd.Flags().String("page-id", "", "The Amazon Resource Name (ARN) of the engagement to a specific contact channel.")
		ssmContacts_listPageReceiptsCmd.MarkFlagRequired("page-id")
	})
	ssmContactsCmd.AddCommand(ssmContacts_listPageReceiptsCmd)
}
