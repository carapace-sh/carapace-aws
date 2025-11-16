package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var odb_createCloudExadataInfrastructureCmd = &cobra.Command{
	Use:   "create-cloud-exadata-infrastructure",
	Short: "Creates an Exadata infrastructure.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(odb_createCloudExadataInfrastructureCmd).Standalone()

	odb_createCloudExadataInfrastructureCmd.Flags().String("availability-zone", "", "The name of the Availability Zone (AZ) where the Exadata infrastructure is located.")
	odb_createCloudExadataInfrastructureCmd.Flags().String("availability-zone-id", "", "The AZ ID of the AZ where the Exadata infrastructure is located.")
	odb_createCloudExadataInfrastructureCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	odb_createCloudExadataInfrastructureCmd.Flags().String("compute-count", "", "The number of database servers for the Exadata infrastructure.")
	odb_createCloudExadataInfrastructureCmd.Flags().String("customer-contacts-to-send-to-oci", "", "The email addresses of contacts to receive notification from Oracle about maintenance updates for the Exadata infrastructure.")
	odb_createCloudExadataInfrastructureCmd.Flags().String("database-server-type", "", "The database server model type of the Exadata infrastructure.")
	odb_createCloudExadataInfrastructureCmd.Flags().String("display-name", "", "A user-friendly name for the Exadata infrastructure.")
	odb_createCloudExadataInfrastructureCmd.Flags().String("maintenance-window", "", "The maintenance window configuration for the Exadata Cloud infrastructure.")
	odb_createCloudExadataInfrastructureCmd.Flags().String("shape", "", "The model name of the Exadata infrastructure.")
	odb_createCloudExadataInfrastructureCmd.Flags().String("storage-count", "", "The number of storage servers to activate for this Exadata infrastructure.")
	odb_createCloudExadataInfrastructureCmd.Flags().String("storage-server-type", "", "The storage server model type of the Exadata infrastructure.")
	odb_createCloudExadataInfrastructureCmd.Flags().String("tags", "", "The list of resource tags to apply to the Exadata infrastructure.")
	odb_createCloudExadataInfrastructureCmd.MarkFlagRequired("compute-count")
	odb_createCloudExadataInfrastructureCmd.MarkFlagRequired("display-name")
	odb_createCloudExadataInfrastructureCmd.MarkFlagRequired("shape")
	odb_createCloudExadataInfrastructureCmd.MarkFlagRequired("storage-count")
	odbCmd.AddCommand(odb_createCloudExadataInfrastructureCmd)
}
