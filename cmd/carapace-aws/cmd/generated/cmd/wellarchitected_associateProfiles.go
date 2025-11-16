package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wellarchitected_associateProfilesCmd = &cobra.Command{
	Use:   "associate-profiles",
	Short: "Associate a profile with a workload.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wellarchitected_associateProfilesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(wellarchitected_associateProfilesCmd).Standalone()

		wellarchitected_associateProfilesCmd.Flags().String("profile-arns", "", "The list of profile ARNs to associate with the workload.")
		wellarchitected_associateProfilesCmd.Flags().String("workload-id", "", "")
		wellarchitected_associateProfilesCmd.MarkFlagRequired("profile-arns")
		wellarchitected_associateProfilesCmd.MarkFlagRequired("workload-id")
	})
	wellarchitectedCmd.AddCommand(wellarchitected_associateProfilesCmd)
}
