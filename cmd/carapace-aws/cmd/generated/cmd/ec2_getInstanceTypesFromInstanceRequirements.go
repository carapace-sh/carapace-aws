package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_getInstanceTypesFromInstanceRequirementsCmd = &cobra.Command{
	Use:   "get-instance-types-from-instance-requirements",
	Short: "Returns a list of instance types with the specified instance attributes.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_getInstanceTypesFromInstanceRequirementsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_getInstanceTypesFromInstanceRequirementsCmd).Standalone()

		ec2_getInstanceTypesFromInstanceRequirementsCmd.Flags().String("architecture-types", "", "The processor architecture type.")
		ec2_getInstanceTypesFromInstanceRequirementsCmd.Flags().String("context", "", "Reserved.")
		ec2_getInstanceTypesFromInstanceRequirementsCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_getInstanceTypesFromInstanceRequirementsCmd.Flags().String("instance-requirements", "", "The attributes required for the instance types.")
		ec2_getInstanceTypesFromInstanceRequirementsCmd.Flags().String("max-results", "", "The maximum number of items to return for this request.")
		ec2_getInstanceTypesFromInstanceRequirementsCmd.Flags().String("next-token", "", "The token returned from a previous paginated request.")
		ec2_getInstanceTypesFromInstanceRequirementsCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_getInstanceTypesFromInstanceRequirementsCmd.Flags().String("virtualization-types", "", "The virtualization type.")
		ec2_getInstanceTypesFromInstanceRequirementsCmd.MarkFlagRequired("architecture-types")
		ec2_getInstanceTypesFromInstanceRequirementsCmd.MarkFlagRequired("instance-requirements")
		ec2_getInstanceTypesFromInstanceRequirementsCmd.Flag("no-dry-run").Hidden = true
		ec2_getInstanceTypesFromInstanceRequirementsCmd.MarkFlagRequired("virtualization-types")
	})
	ec2Cmd.AddCommand(ec2_getInstanceTypesFromInstanceRequirementsCmd)
}
