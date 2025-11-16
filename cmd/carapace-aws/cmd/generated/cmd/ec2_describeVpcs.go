package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeVpcsCmd = &cobra.Command{
	Use:   "describe-vpcs",
	Short: "Describes your VPCs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeVpcsCmd).Standalone()

	ec2_describeVpcsCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeVpcsCmd.Flags().String("filters", "", "The filters.")
	ec2_describeVpcsCmd.Flags().String("max-results", "", "The maximum number of items to return for this request.")
	ec2_describeVpcsCmd.Flags().String("next-token", "", "The token returned from a previous paginated request.")
	ec2_describeVpcsCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeVpcsCmd.Flags().String("vpc-ids", "", "The IDs of the VPCs.")
	ec2_describeVpcsCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_describeVpcsCmd)
}
