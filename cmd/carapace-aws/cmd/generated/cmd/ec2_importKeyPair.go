package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_importKeyPairCmd = &cobra.Command{
	Use:   "import-key-pair",
	Short: "Imports the public key from an RSA or ED25519 key pair that you created using a third-party tool.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_importKeyPairCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_importKeyPairCmd).Standalone()

		ec2_importKeyPairCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_importKeyPairCmd.Flags().String("key-name", "", "A unique name for the key pair.")
		ec2_importKeyPairCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_importKeyPairCmd.Flags().String("public-key-material", "", "The public key.")
		ec2_importKeyPairCmd.Flags().String("tag-specifications", "", "The tags to apply to the imported key pair.")
		ec2_importKeyPairCmd.MarkFlagRequired("key-name")
		ec2_importKeyPairCmd.Flag("no-dry-run").Hidden = true
		ec2_importKeyPairCmd.MarkFlagRequired("public-key-material")
	})
	ec2Cmd.AddCommand(ec2_importKeyPairCmd)
}
