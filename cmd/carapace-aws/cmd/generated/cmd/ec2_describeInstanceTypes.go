package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeInstanceTypesCmd = &cobra.Command{
	Use:   "describe-instance-types",
	Short: "Describes the specified instance types.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeInstanceTypesCmd).Standalone()

	ec2_describeInstanceTypesCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeInstanceTypesCmd.Flags().String("filters", "", "One or more filters.")
	ec2_describeInstanceTypesCmd.Flags().String("instance-types", "", "The instance types.")
	ec2_describeInstanceTypesCmd.Flags().String("max-results", "", "The maximum number of items to return for this request.")
	ec2_describeInstanceTypesCmd.Flags().String("next-token", "", "The token returned from a previous paginated request.")
	ec2_describeInstanceTypesCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeInstanceTypesCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_describeInstanceTypesCmd)
}
