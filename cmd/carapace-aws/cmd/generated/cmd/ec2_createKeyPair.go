package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_createKeyPairCmd = &cobra.Command{
	Use:   "create-key-pair",
	Short: "Creates an ED25519 or 2048-bit RSA key pair with the specified name and in the specified format.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_createKeyPairCmd).Standalone()

	ec2_createKeyPairCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_createKeyPairCmd.Flags().String("key-format", "", "The format of the key pair.")
	ec2_createKeyPairCmd.Flags().String("key-name", "", "A unique name for the key pair.")
	ec2_createKeyPairCmd.Flags().String("key-type", "", "The type of key pair.")
	ec2_createKeyPairCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_createKeyPairCmd.Flags().String("tag-specifications", "", "The tags to apply to the new key pair.")
	ec2_createKeyPairCmd.MarkFlagRequired("key-name")
	ec2_createKeyPairCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_createKeyPairCmd)
}
