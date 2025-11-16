package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wellarchitected_updateProfileCmd = &cobra.Command{
	Use:   "update-profile",
	Short: "Update a profile.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wellarchitected_updateProfileCmd).Standalone()

	wellarchitected_updateProfileCmd.Flags().String("profile-arn", "", "The profile ARN.")
	wellarchitected_updateProfileCmd.Flags().String("profile-description", "", "The profile description.")
	wellarchitected_updateProfileCmd.Flags().String("profile-questions", "", "Profile questions.")
	wellarchitected_updateProfileCmd.MarkFlagRequired("profile-arn")
	wellarchitectedCmd.AddCommand(wellarchitected_updateProfileCmd)
}
