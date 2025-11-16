package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_disableImageBlockPublicAccessCmd = &cobra.Command{
	Use:   "disable-image-block-public-access",
	Short: "Disables *block public access for AMIs* at the account level in the specified Amazon Web Services Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_disableImageBlockPublicAccessCmd).Standalone()

	ec2_disableImageBlockPublicAccessCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_disableImageBlockPublicAccessCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_disableImageBlockPublicAccessCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_disableImageBlockPublicAccessCmd)
}
