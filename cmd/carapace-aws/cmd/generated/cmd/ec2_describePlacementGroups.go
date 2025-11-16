package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describePlacementGroupsCmd = &cobra.Command{
	Use:   "describe-placement-groups",
	Short: "Describes the specified placement groups or all of your placement groups.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describePlacementGroupsCmd).Standalone()

	ec2_describePlacementGroupsCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the operation, without actually making the request, and provides an error response.")
	ec2_describePlacementGroupsCmd.Flags().String("filters", "", "The filters.")
	ec2_describePlacementGroupsCmd.Flags().String("group-ids", "", "The IDs of the placement groups.")
	ec2_describePlacementGroupsCmd.Flags().String("group-names", "", "The names of the placement groups.")
	ec2_describePlacementGroupsCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the operation, without actually making the request, and provides an error response.")
	ec2_describePlacementGroupsCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_describePlacementGroupsCmd)
}
