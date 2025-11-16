package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var transfer_createAccessCmd = &cobra.Command{
	Use:   "create-access",
	Short: "Used by administrators to choose which groups in the directory should have access to upload and download files over the enabled protocols using Transfer Family.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(transfer_createAccessCmd).Standalone()

	transfer_createAccessCmd.Flags().String("external-id", "", "A unique identifier that is required to identify specific groups within your directory.")
	transfer_createAccessCmd.Flags().String("home-directory", "", "The landing directory (folder) for a user when they log in to the server using the client.")
	transfer_createAccessCmd.Flags().String("home-directory-mappings", "", "Logical directory mappings that specify what Amazon S3 or Amazon EFS paths and keys should be visible to your user and how you want to make them visible.")
	transfer_createAccessCmd.Flags().String("home-directory-type", "", "The type of landing directory (folder) that you want your users' home directory to be when they log in to the server.")
	transfer_createAccessCmd.Flags().String("policy", "", "A session policy for your user so that you can use the same Identity and Access Management (IAM) role across multiple users.")
	transfer_createAccessCmd.Flags().String("posix-profile", "", "")
	transfer_createAccessCmd.Flags().String("role", "", "The Amazon Resource Name (ARN) of the Identity and Access Management (IAM) role that controls your users' access to your Amazon S3 bucket or Amazon EFS file system.")
	transfer_createAccessCmd.Flags().String("server-id", "", "A system-assigned unique identifier for a server instance.")
	transfer_createAccessCmd.MarkFlagRequired("external-id")
	transfer_createAccessCmd.MarkFlagRequired("role")
	transfer_createAccessCmd.MarkFlagRequired("server-id")
	transferCmd.AddCommand(transfer_createAccessCmd)
}
