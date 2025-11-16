package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeTrunkInterfaceAssociationsCmd = &cobra.Command{
	Use:   "describe-trunk-interface-associations",
	Short: "Describes one or more network interface trunk associations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeTrunkInterfaceAssociationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_describeTrunkInterfaceAssociationsCmd).Standalone()

		ec2_describeTrunkInterfaceAssociationsCmd.Flags().String("association-ids", "", "The IDs of the associations.")
		ec2_describeTrunkInterfaceAssociationsCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeTrunkInterfaceAssociationsCmd.Flags().String("filters", "", "One or more filters.")
		ec2_describeTrunkInterfaceAssociationsCmd.Flags().String("max-results", "", "The maximum number of results to return with a single call.")
		ec2_describeTrunkInterfaceAssociationsCmd.Flags().String("next-token", "", "The token for the next page of results.")
		ec2_describeTrunkInterfaceAssociationsCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeTrunkInterfaceAssociationsCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_describeTrunkInterfaceAssociationsCmd)
}
