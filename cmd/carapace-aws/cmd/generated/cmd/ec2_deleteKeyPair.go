package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_deleteKeyPairCmd = &cobra.Command{
	Use:   "delete-key-pair",
	Short: "Deletes the specified key pair, by removing the public key from Amazon EC2.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_deleteKeyPairCmd).Standalone()

	ec2_deleteKeyPairCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_deleteKeyPairCmd.Flags().String("key-name", "", "The name of the key pair.")
	ec2_deleteKeyPairCmd.Flags().String("key-pair-id", "", "The ID of the key pair.")
	ec2_deleteKeyPairCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_deleteKeyPairCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_deleteKeyPairCmd)
}
