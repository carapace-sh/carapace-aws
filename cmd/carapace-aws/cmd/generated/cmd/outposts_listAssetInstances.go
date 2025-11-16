package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var outposts_listAssetInstancesCmd = &cobra.Command{
	Use:   "list-asset-instances",
	Short: "A list of Amazon EC2 instances, belonging to all accounts, running on the specified Outpost.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(outposts_listAssetInstancesCmd).Standalone()

	outposts_listAssetInstancesCmd.Flags().String("account-id-filter", "", "Filters the results by account ID.")
	outposts_listAssetInstancesCmd.Flags().String("asset-id-filter", "", "Filters the results by asset ID.")
	outposts_listAssetInstancesCmd.Flags().String("aws-service-filter", "", "Filters the results by Amazon Web Services service.")
	outposts_listAssetInstancesCmd.Flags().String("instance-type-filter", "", "Filters the results by instance ID.")
	outposts_listAssetInstancesCmd.Flags().String("max-results", "", "")
	outposts_listAssetInstancesCmd.Flags().String("next-token", "", "")
	outposts_listAssetInstancesCmd.Flags().String("outpost-identifier", "", "The ID of the Outpost.")
	outposts_listAssetInstancesCmd.MarkFlagRequired("outpost-identifier")
	outpostsCmd.AddCommand(outposts_listAssetInstancesCmd)
}
