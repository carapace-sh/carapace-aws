package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fsx_describeDataRepositoryAssociationsCmd = &cobra.Command{
	Use:   "describe-data-repository-associations",
	Short: "Returns the description of specific Amazon FSx for Lustre or Amazon File Cache data repository associations, if one or more `AssociationIds` values are provided in the request, or if filters are used in the request.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fsx_describeDataRepositoryAssociationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(fsx_describeDataRepositoryAssociationsCmd).Standalone()

		fsx_describeDataRepositoryAssociationsCmd.Flags().String("association-ids", "", "IDs of the data repository associations whose descriptions you want to retrieve (String).")
		fsx_describeDataRepositoryAssociationsCmd.Flags().String("filters", "", "")
		fsx_describeDataRepositoryAssociationsCmd.Flags().String("max-results", "", "The maximum number of resources to return in the response.")
		fsx_describeDataRepositoryAssociationsCmd.Flags().String("next-token", "", "")
	})
	fsxCmd.AddCommand(fsx_describeDataRepositoryAssociationsCmd)
}
