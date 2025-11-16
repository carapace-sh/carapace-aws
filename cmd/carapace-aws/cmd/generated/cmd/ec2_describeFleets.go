package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeFleetsCmd = &cobra.Command{
	Use:   "describe-fleets",
	Short: "Describes the specified EC2 Fleet or all of your EC2 Fleets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeFleetsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_describeFleetsCmd).Standalone()

		ec2_describeFleetsCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeFleetsCmd.Flags().String("filters", "", "The filters.")
		ec2_describeFleetsCmd.Flags().String("fleet-ids", "", "The IDs of the EC2 Fleets.")
		ec2_describeFleetsCmd.Flags().String("max-results", "", "The maximum number of items to return for this request.")
		ec2_describeFleetsCmd.Flags().String("next-token", "", "The token returned from a previous paginated request.")
		ec2_describeFleetsCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeFleetsCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_describeFleetsCmd)
}
