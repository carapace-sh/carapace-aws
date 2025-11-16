package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var odb_listSystemVersionsCmd = &cobra.Command{
	Use:   "list-system-versions",
	Short: "Returns information about the system versions that are available for a VM cluster for the specified `giVersion` and `shape`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(odb_listSystemVersionsCmd).Standalone()

	odb_listSystemVersionsCmd.Flags().String("gi-version", "", "The software version of the Exadata Grid Infrastructure (GI).")
	odb_listSystemVersionsCmd.Flags().String("max-results", "", "The maximum number of items to return for this request.")
	odb_listSystemVersionsCmd.Flags().String("next-token", "", "The token returned from a previous paginated request.")
	odb_listSystemVersionsCmd.Flags().String("shape", "", "The Exadata hardware system model.")
	odb_listSystemVersionsCmd.MarkFlagRequired("gi-version")
	odb_listSystemVersionsCmd.MarkFlagRequired("shape")
	odbCmd.AddCommand(odb_listSystemVersionsCmd)
}
