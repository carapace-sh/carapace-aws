package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var odb_listGiVersionsCmd = &cobra.Command{
	Use:   "list-gi-versions",
	Short: "Returns information about Oracle Grid Infrastructure (GI) software versions that are available for a VM cluster for the specified shape.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(odb_listGiVersionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(odb_listGiVersionsCmd).Standalone()

		odb_listGiVersionsCmd.Flags().String("max-results", "", "The maximum number of items to return for this request.")
		odb_listGiVersionsCmd.Flags().String("next-token", "", "The token returned from a previous paginated request.")
		odb_listGiVersionsCmd.Flags().String("shape", "", "The shape to return GI versions for.")
	})
	odbCmd.AddCommand(odb_listGiVersionsCmd)
}
