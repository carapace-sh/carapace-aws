package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_describePatchGroupsCmd = &cobra.Command{
	Use:   "describe-patch-groups",
	Short: "Lists all patch groups that have been registered with patch baselines.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_describePatchGroupsCmd).Standalone()

	ssm_describePatchGroupsCmd.Flags().String("filters", "", "Each element in the array is a structure containing a key-value pair.")
	ssm_describePatchGroupsCmd.Flags().String("max-results", "", "The maximum number of patch groups to return (per page).")
	ssm_describePatchGroupsCmd.Flags().String("next-token", "", "The token for the next set of items to return.")
	ssmCmd.AddCommand(ssm_describePatchGroupsCmd)
}
