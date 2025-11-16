package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var transfer_updateAccessCmd = &cobra.Command{
	Use:   "update-access",
	Short: "Allows you to update parameters for the access specified in the `ServerID` and `ExternalID` parameters.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(transfer_updateAccessCmd).Standalone()

	transfer_updateAccessCmd.Flags().String("external-id", "", "A unique identifier that is required to identify specific groups within your directory.")
	transfer_updateAccessCmd.Flags().String("home-directory", "", "The landing directory (folder) for a user when they log in to the server using the client.")
	transfer_updateAccessCmd.Flags().String("home-directory-mappings", "", "Logical directory mappings that specify what Amazon S3 or Amazon EFS paths and keys should be visible to your user and how you want to make them visible.")
	transfer_updateAccessCmd.Flags().String("home-directory-type", "", "The type of landing directory (folder) that you want your users' home directory to be when they log in to the server.")
	transfer_updateAccessCmd.Flags().String("policy", "", "A session policy for your user so that you can use the same Identity and Access Management (IAM) role across multiple users.")
	transfer_updateAccessCmd.Flags().String("posix-profile", "", "")
	transfer_updateAccessCmd.Flags().String("role", "", "The Amazon Resource Name (ARN) of the Identity and Access Management (IAM) role that controls your users' access to your Amazon S3 bucket or Amazon EFS file system.")
	transfer_updateAccessCmd.Flags().String("server-id", "", "A system-assigned unique identifier for a server instance.")
	transfer_updateAccessCmd.MarkFlagRequired("external-id")
	transfer_updateAccessCmd.MarkFlagRequired("server-id")
	transferCmd.AddCommand(transfer_updateAccessCmd)
}
