package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var odb_deleteCloudExadataInfrastructureCmd = &cobra.Command{
	Use:   "delete-cloud-exadata-infrastructure",
	Short: "Deletes the specified Exadata infrastructure.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(odb_deleteCloudExadataInfrastructureCmd).Standalone()

	odb_deleteCloudExadataInfrastructureCmd.Flags().String("cloud-exadata-infrastructure-id", "", "The unique identifier of the Exadata infrastructure to delete.")
	odb_deleteCloudExadataInfrastructureCmd.MarkFlagRequired("cloud-exadata-infrastructure-id")
	odbCmd.AddCommand(odb_deleteCloudExadataInfrastructureCmd)
}
