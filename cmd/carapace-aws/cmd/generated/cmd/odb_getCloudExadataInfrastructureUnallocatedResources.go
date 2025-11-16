package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var odb_getCloudExadataInfrastructureUnallocatedResourcesCmd = &cobra.Command{
	Use:   "get-cloud-exadata-infrastructure-unallocated-resources",
	Short: "Retrieves information about unallocated resources in a specified Cloud Exadata Infrastructure.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(odb_getCloudExadataInfrastructureUnallocatedResourcesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(odb_getCloudExadataInfrastructureUnallocatedResourcesCmd).Standalone()

		odb_getCloudExadataInfrastructureUnallocatedResourcesCmd.Flags().String("cloud-exadata-infrastructure-id", "", "The unique identifier of the Cloud Exadata infrastructure for which to retrieve unallocated resources.")
		odb_getCloudExadataInfrastructureUnallocatedResourcesCmd.Flags().String("db-servers", "", "The database servers to include in the unallocated resources query.")
		odb_getCloudExadataInfrastructureUnallocatedResourcesCmd.MarkFlagRequired("cloud-exadata-infrastructure-id")
	})
	odbCmd.AddCommand(odb_getCloudExadataInfrastructureUnallocatedResourcesCmd)
}
