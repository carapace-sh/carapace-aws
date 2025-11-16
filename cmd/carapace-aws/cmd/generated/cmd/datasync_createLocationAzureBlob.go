package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datasync_createLocationAzureBlobCmd = &cobra.Command{
	Use:   "create-location-azure-blob",
	Short: "Creates a transfer *location* for a Microsoft Azure Blob Storage container.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datasync_createLocationAzureBlobCmd).Standalone()

	datasync_createLocationAzureBlobCmd.Flags().String("access-tier", "", "Specifies the access tier that you want your objects or files transferred into.")
	datasync_createLocationAzureBlobCmd.Flags().String("agent-arns", "", "(Optional) Specifies the Amazon Resource Name (ARN) of the DataSync agent that can connect with your Azure Blob Storage container.")
	datasync_createLocationAzureBlobCmd.Flags().String("authentication-type", "", "Specifies the authentication method DataSync uses to access your Azure Blob Storage.")
	datasync_createLocationAzureBlobCmd.Flags().String("blob-type", "", "Specifies the type of blob that you want your objects or files to be when transferring them into Azure Blob Storage.")
	datasync_createLocationAzureBlobCmd.Flags().String("cmk-secret-config", "", "Specifies configuration information for a DataSync-managed secret, which includes the authentication token that DataSync uses to access a specific AzureBlob storage location, with a customer-managed KMS key.")
	datasync_createLocationAzureBlobCmd.Flags().String("container-url", "", "Specifies the URL of the Azure Blob Storage container involved in your transfer.")
	datasync_createLocationAzureBlobCmd.Flags().String("custom-secret-config", "", "Specifies configuration information for a customer-managed Secrets Manager secret where the authentication token for an AzureBlob storage location is stored in plain text.")
	datasync_createLocationAzureBlobCmd.Flags().String("sas-configuration", "", "Specifies the SAS configuration that allows DataSync to access your Azure Blob Storage.")
	datasync_createLocationAzureBlobCmd.Flags().String("subdirectory", "", "Specifies path segments if you want to limit your transfer to a virtual directory in your container (for example, `/my/images`).")
	datasync_createLocationAzureBlobCmd.Flags().String("tags", "", "Specifies labels that help you categorize, filter, and search for your Amazon Web Services resources.")
	datasync_createLocationAzureBlobCmd.MarkFlagRequired("authentication-type")
	datasync_createLocationAzureBlobCmd.MarkFlagRequired("container-url")
	datasyncCmd.AddCommand(datasync_createLocationAzureBlobCmd)
}
