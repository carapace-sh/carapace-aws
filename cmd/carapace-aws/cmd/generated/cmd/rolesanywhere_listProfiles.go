package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rolesanywhere_listProfilesCmd = &cobra.Command{
	Use:   "list-profiles",
	Short: "Lists all profiles in the authenticated account and Amazon Web Services Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rolesanywhere_listProfilesCmd).Standalone()

	rolesanywhere_listProfilesCmd.Flags().String("next-token", "", "A token that indicates where the output should continue from, if a previous request did not show all results.")
	rolesanywhere_listProfilesCmd.Flags().String("page-size", "", "The number of resources in the paginated list.")
	rolesanywhereCmd.AddCommand(rolesanywhere_listProfilesCmd)
}
