package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wellarchitected_disassociateProfilesCmd = &cobra.Command{
	Use:   "disassociate-profiles",
	Short: "Disassociate a profile from a workload.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wellarchitected_disassociateProfilesCmd).Standalone()

	wellarchitected_disassociateProfilesCmd.Flags().String("profile-arns", "", "The list of profile ARNs to disassociate from the workload.")
	wellarchitected_disassociateProfilesCmd.Flags().String("workload-id", "", "")
	wellarchitected_disassociateProfilesCmd.MarkFlagRequired("profile-arns")
	wellarchitected_disassociateProfilesCmd.MarkFlagRequired("workload-id")
	wellarchitectedCmd.AddCommand(wellarchitected_disassociateProfilesCmd)
}
