package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeAccountAttributesCmd = &cobra.Command{
	Use:   "describe-account-attributes",
	Short: "Describes attributes of your Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeAccountAttributesCmd).Standalone()

	ec2_describeAccountAttributesCmd.Flags().String("attribute-names", "", "The account attribute names.")
	ec2_describeAccountAttributesCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeAccountAttributesCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeAccountAttributesCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_describeAccountAttributesCmd)
}
