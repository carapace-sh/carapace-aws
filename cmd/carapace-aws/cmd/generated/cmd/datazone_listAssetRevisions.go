package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_listAssetRevisionsCmd = &cobra.Command{
	Use:   "list-asset-revisions",
	Short: "Lists the revisions for the asset.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_listAssetRevisionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(datazone_listAssetRevisionsCmd).Standalone()

		datazone_listAssetRevisionsCmd.Flags().String("domain-identifier", "", "The identifier of the domain.")
		datazone_listAssetRevisionsCmd.Flags().String("identifier", "", "The identifier of the asset.")
		datazone_listAssetRevisionsCmd.Flags().String("max-results", "", "The maximum number of revisions to return in a single call to `ListAssetRevisions`.")
		datazone_listAssetRevisionsCmd.Flags().String("next-token", "", "When the number of revisions is greater than the default value for the `MaxResults` parameter, or if you explicitly specify a value for `MaxResults` that is less than the number of revisions, the response includes a pagination token named `NextToken`.")
		datazone_listAssetRevisionsCmd.MarkFlagRequired("domain-identifier")
		datazone_listAssetRevisionsCmd.MarkFlagRequired("identifier")
	})
	datazoneCmd.AddCommand(datazone_listAssetRevisionsCmd)
}
