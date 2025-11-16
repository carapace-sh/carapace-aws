package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeInstanceTopologyCmd = &cobra.Command{
	Use:   "describe-instance-topology",
	Short: "Describes a tree-based hierarchy that represents the physical host placement of your EC2 instances within an Availability Zone or Local Zone.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeInstanceTopologyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_describeInstanceTopologyCmd).Standalone()

		ec2_describeInstanceTopologyCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the operation, without actually making the request, and provides an error response.")
		ec2_describeInstanceTopologyCmd.Flags().String("filters", "", "The filters.")
		ec2_describeInstanceTopologyCmd.Flags().String("group-names", "", "The name of the placement group that each instance is in.")
		ec2_describeInstanceTopologyCmd.Flags().String("instance-ids", "", "The instance IDs.")
		ec2_describeInstanceTopologyCmd.Flags().String("max-results", "", "The maximum number of items to return for this request.")
		ec2_describeInstanceTopologyCmd.Flags().String("next-token", "", "The token returned from a previous paginated request.")
		ec2_describeInstanceTopologyCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the operation, without actually making the request, and provides an error response.")
		ec2_describeInstanceTopologyCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_describeInstanceTopologyCmd)
}
