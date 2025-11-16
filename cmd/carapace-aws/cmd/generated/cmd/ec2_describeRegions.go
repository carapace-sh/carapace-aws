package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeRegionsCmd = &cobra.Command{
	Use:   "describe-regions",
	Short: "Describes the Regions that are enabled for your account, or all Regions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeRegionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_describeRegionsCmd).Standalone()

		ec2_describeRegionsCmd.Flags().Bool("all-regions", false, "Indicates whether to display all Regions, including Regions that are disabled for your account.")
		ec2_describeRegionsCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeRegionsCmd.Flags().String("filters", "", "The filters.")
		ec2_describeRegionsCmd.Flags().Bool("no-all-regions", false, "Indicates whether to display all Regions, including Regions that are disabled for your account.")
		ec2_describeRegionsCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeRegionsCmd.Flags().String("region-names", "", "The names of the Regions.")
		ec2_describeRegionsCmd.Flag("no-all-regions").Hidden = true
		ec2_describeRegionsCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_describeRegionsCmd)
}
