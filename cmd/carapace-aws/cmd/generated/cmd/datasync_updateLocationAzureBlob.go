package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datasync_updateLocationAzureBlobCmd = &cobra.Command{
	Use:   "update-location-azure-blob",
	Short: "Modifies the following configurations of the Microsoft Azure Blob Storage transfer location that you're using with DataSync.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datasync_updateLocationAzureBlobCmd).Standalone()

	datasync_updateLocationAzureBlobCmd.Flags().String("access-tier", "", "Specifies the access tier that you want your objects or files transferred into.")
	datasync_updateLocationAzureBlobCmd.Flags().String("agent-arns", "", "(Optional) Specifies the Amazon Resource Name (ARN) of the DataSync agent that can connect with your Azure Blob Storage container.")
	datasync_updateLocationAzureBlobCmd.Flags().String("authentication-type", "", "Specifies the authentication method DataSync uses to access your Azure Blob Storage.")
	datasync_updateLocationAzureBlobCmd.Flags().String("blob-type", "", "Specifies the type of blob that you want your objects or files to be when transferring them into Azure Blob Storage.")
	datasync_updateLocationAzureBlobCmd.Flags().String("cmk-secret-config", "", "Specifies configuration information for a DataSync-managed secret, such as an authentication token or set of credentials that DataSync uses to access a specific transfer location, and a customer-managed KMS key.")
	datasync_updateLocationAzureBlobCmd.Flags().String("custom-secret-config", "", "Specifies configuration information for a customer-managed secret, such as an authentication token or set of credentials that DataSync uses to access a specific transfer location, and a customer-managed KMS key.")
	datasync_updateLocationAzureBlobCmd.Flags().String("location-arn", "", "Specifies the ARN of the Azure Blob Storage transfer location that you're updating.")
	datasync_updateLocationAzureBlobCmd.Flags().String("sas-configuration", "", "Specifies the SAS configuration that allows DataSync to access your Azure Blob Storage.")
	datasync_updateLocationAzureBlobCmd.Flags().String("subdirectory", "", "Specifies path segments if you want to limit your transfer to a virtual directory in your container (for example, `/my/images`).")
	datasync_updateLocationAzureBlobCmd.MarkFlagRequired("location-arn")
	datasyncCmd.AddCommand(datasync_updateLocationAzureBlobCmd)
}
