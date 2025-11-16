package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var synthetics_listGroupsCmd = &cobra.Command{
	Use:   "list-groups",
	Short: "Returns a list of all groups in the account, displaying their names, unique IDs, and ARNs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(synthetics_listGroupsCmd).Standalone()

	synthetics_listGroupsCmd.Flags().String("max-results", "", "Specify this parameter to limit how many groups are returned each time you use the `ListGroups` operation.")
	synthetics_listGroupsCmd.Flags().String("next-token", "", "A token that indicates that there is more data available.")
	syntheticsCmd.AddCommand(synthetics_listGroupsCmd)
}
