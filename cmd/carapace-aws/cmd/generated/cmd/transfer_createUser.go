package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var transfer_createUserCmd = &cobra.Command{
	Use:   "create-user",
	Short: "Creates a user and associates them with an existing file transfer protocol-enabled server.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(transfer_createUserCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(transfer_createUserCmd).Standalone()

		transfer_createUserCmd.Flags().String("home-directory", "", "The landing directory (folder) for a user when they log in to the server using the client.")
		transfer_createUserCmd.Flags().String("home-directory-mappings", "", "Logical directory mappings that specify what Amazon S3 or Amazon EFS paths and keys should be visible to your user and how you want to make them visible.")
		transfer_createUserCmd.Flags().String("home-directory-type", "", "The type of landing directory (folder) that you want your users' home directory to be when they log in to the server.")
		transfer_createUserCmd.Flags().String("policy", "", "A session policy for your user so that you can use the same Identity and Access Management (IAM) role across multiple users.")
		transfer_createUserCmd.Flags().String("posix-profile", "", "Specifies the full POSIX identity, including user ID (`Uid`), group ID (`Gid`), and any secondary groups IDs (`SecondaryGids`), that controls your users' access to your Amazon EFS file systems.")
		transfer_createUserCmd.Flags().String("role", "", "The Amazon Resource Name (ARN) of the Identity and Access Management (IAM) role that controls your users' access to your Amazon S3 bucket or Amazon EFS file system.")
		transfer_createUserCmd.Flags().String("server-id", "", "A system-assigned unique identifier for a server instance.")
		transfer_createUserCmd.Flags().String("ssh-public-key-body", "", "The public portion of the Secure Shell (SSH) key used to authenticate the user to the server.")
		transfer_createUserCmd.Flags().String("tags", "", "Key-value pairs that can be used to group and search for users.")
		transfer_createUserCmd.Flags().String("user-name", "", "A unique string that identifies a user and is associated with a `ServerId`.")
		transfer_createUserCmd.MarkFlagRequired("role")
		transfer_createUserCmd.MarkFlagRequired("server-id")
		transfer_createUserCmd.MarkFlagRequired("user-name")
	})
	transferCmd.AddCommand(transfer_createUserCmd)
}
