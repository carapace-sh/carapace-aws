package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lakeformation_listLakeFormationOptInsCmd = &cobra.Command{
	Use:   "list-lake-formation-opt-ins",
	Short: "Retrieve the current list of resources and principals that are opt in to enforce Lake Formation permissions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lakeformation_listLakeFormationOptInsCmd).Standalone()

	lakeformation_listLakeFormationOptInsCmd.Flags().String("max-results", "", "The maximum number of results to return.")
	lakeformation_listLakeFormationOptInsCmd.Flags().String("next-token", "", "A continuation token, if this is not the first call to retrieve this list.")
	lakeformation_listLakeFormationOptInsCmd.Flags().String("principal", "", "")
	lakeformation_listLakeFormationOptInsCmd.Flags().String("resource", "", "A structure for the resource.")
	lakeformationCmd.AddCommand(lakeformation_listLakeFormationOptInsCmd)
}
