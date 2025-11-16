package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var odb_listCloudExadataInfrastructuresCmd = &cobra.Command{
	Use:   "list-cloud-exadata-infrastructures",
	Short: "Returns information about the Exadata infrastructures owned by your Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(odb_listCloudExadataInfrastructuresCmd).Standalone()

	odb_listCloudExadataInfrastructuresCmd.Flags().String("max-results", "", "The maximum number of items to return for this request.")
	odb_listCloudExadataInfrastructuresCmd.Flags().String("next-token", "", "The token returned from a previous paginated request.")
	odbCmd.AddCommand(odb_listCloudExadataInfrastructuresCmd)
}
