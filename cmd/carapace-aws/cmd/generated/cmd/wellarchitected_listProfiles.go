package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wellarchitected_listProfilesCmd = &cobra.Command{
	Use:   "list-profiles",
	Short: "List profiles.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wellarchitected_listProfilesCmd).Standalone()

	wellarchitected_listProfilesCmd.Flags().String("max-results", "", "")
	wellarchitected_listProfilesCmd.Flags().String("next-token", "", "")
	wellarchitected_listProfilesCmd.Flags().String("profile-name-prefix", "", "An optional string added to the beginning of each profile name returned in the results.")
	wellarchitected_listProfilesCmd.Flags().String("profile-owner-type", "", "Profile owner type.")
	wellarchitectedCmd.AddCommand(wellarchitected_listProfilesCmd)
}
