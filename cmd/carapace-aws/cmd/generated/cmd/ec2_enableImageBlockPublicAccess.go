package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_enableImageBlockPublicAccessCmd = &cobra.Command{
	Use:   "enable-image-block-public-access",
	Short: "Enables *block public access for AMIs* at the account level in the specified Amazon Web Services Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_enableImageBlockPublicAccessCmd).Standalone()

	ec2_enableImageBlockPublicAccessCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_enableImageBlockPublicAccessCmd.Flags().String("image-block-public-access-state", "", "Specify `block-new-sharing` to enable block public access for AMIs at the account level in the specified Region.")
	ec2_enableImageBlockPublicAccessCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_enableImageBlockPublicAccessCmd.MarkFlagRequired("image-block-public-access-state")
	ec2_enableImageBlockPublicAccessCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_enableImageBlockPublicAccessCmd)
}
