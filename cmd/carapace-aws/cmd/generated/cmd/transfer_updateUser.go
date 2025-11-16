package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var transfer_updateUserCmd = &cobra.Command{
	Use:   "update-user",
	Short: "Assigns new properties to a user.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(transfer_updateUserCmd).Standalone()

	transfer_updateUserCmd.Flags().String("home-directory", "", "The landing directory (folder) for a user when they log in to the server using the client.")
	transfer_updateUserCmd.Flags().String("home-directory-mappings", "", "Logical directory mappings that specify what Amazon S3 or Amazon EFS paths and keys should be visible to your user and how you want to make them visible.")
	transfer_updateUserCmd.Flags().String("home-directory-type", "", "The type of landing directory (folder) that you want your users' home directory to be when they log in to the server.")
	transfer_updateUserCmd.Flags().String("policy", "", "A session policy for your user so that you can use the same Identity and Access Management (IAM) role across multiple users.")
	transfer_updateUserCmd.Flags().String("posix-profile", "", "Specifies the full POSIX identity, including user ID (`Uid`), group ID (`Gid`), and any secondary groups IDs (`SecondaryGids`), that controls your users' access to your Amazon Elastic File Systems (Amazon EFS).")
	transfer_updateUserCmd.Flags().String("role", "", "The Amazon Resource Name (ARN) of the Identity and Access Management (IAM) role that controls your users' access to your Amazon S3 bucket or Amazon EFS file system.")
	transfer_updateUserCmd.Flags().String("server-id", "", "A system-assigned unique identifier for a Transfer Family server instance that the user is assigned to.")
	transfer_updateUserCmd.Flags().String("user-name", "", "A unique string that identifies a user and is associated with a server as specified by the `ServerId`.")
	transfer_updateUserCmd.MarkFlagRequired("server-id")
	transfer_updateUserCmd.MarkFlagRequired("user-name")
	transferCmd.AddCommand(transfer_updateUserCmd)
}
