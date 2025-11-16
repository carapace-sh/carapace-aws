package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_getImageBlockPublicAccessStateCmd = &cobra.Command{
	Use:   "get-image-block-public-access-state",
	Short: "Gets the current state of *block public access for AMIs* at the account level in the specified Amazon Web Services Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_getImageBlockPublicAccessStateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_getImageBlockPublicAccessStateCmd).Standalone()

		ec2_getImageBlockPublicAccessStateCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_getImageBlockPublicAccessStateCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_getImageBlockPublicAccessStateCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_getImageBlockPublicAccessStateCmd)
}
