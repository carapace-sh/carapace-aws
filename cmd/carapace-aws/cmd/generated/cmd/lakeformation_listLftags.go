package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lakeformation_listLftagsCmd = &cobra.Command{
	Use:   "list-lftags",
	Short: "Lists LF-tags that the requester has permission to view.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lakeformation_listLftagsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lakeformation_listLftagsCmd).Standalone()

		lakeformation_listLftagsCmd.Flags().String("catalog-id", "", "The identifier for the Data Catalog.")
		lakeformation_listLftagsCmd.Flags().String("max-results", "", "The maximum number of results to return.")
		lakeformation_listLftagsCmd.Flags().String("next-token", "", "A continuation token, if this is not the first call to retrieve this list.")
		lakeformation_listLftagsCmd.Flags().String("resource-share-type", "", "If resource share type is `ALL`, returns both in-account LF-tags and shared LF-tags that the requester has permission to view.")
	})
	lakeformationCmd.AddCommand(lakeformation_listLftagsCmd)
}
