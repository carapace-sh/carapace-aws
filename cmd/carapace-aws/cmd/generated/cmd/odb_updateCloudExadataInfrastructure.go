package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var odb_updateCloudExadataInfrastructureCmd = &cobra.Command{
	Use:   "update-cloud-exadata-infrastructure",
	Short: "Updates the properties of an Exadata infrastructure resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(odb_updateCloudExadataInfrastructureCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(odb_updateCloudExadataInfrastructureCmd).Standalone()

		odb_updateCloudExadataInfrastructureCmd.Flags().String("cloud-exadata-infrastructure-id", "", "The unique identifier of the Exadata infrastructure to update.")
		odb_updateCloudExadataInfrastructureCmd.Flags().String("maintenance-window", "", "")
		odb_updateCloudExadataInfrastructureCmd.MarkFlagRequired("cloud-exadata-infrastructure-id")
	})
	odbCmd.AddCommand(odb_updateCloudExadataInfrastructureCmd)
}
