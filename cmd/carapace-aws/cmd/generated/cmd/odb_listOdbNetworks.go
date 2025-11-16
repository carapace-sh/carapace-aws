package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var odb_listOdbNetworksCmd = &cobra.Command{
	Use:   "list-odb-networks",
	Short: "Returns information about the ODB networks owned by your Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(odb_listOdbNetworksCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(odb_listOdbNetworksCmd).Standalone()

		odb_listOdbNetworksCmd.Flags().String("max-results", "", "The maximum number of items to return for this request.")
		odb_listOdbNetworksCmd.Flags().String("next-token", "", "The token returned from a previous paginated request.")
	})
	odbCmd.AddCommand(odb_listOdbNetworksCmd)
}
