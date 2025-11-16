package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wellarchitected_createProfileCmd = &cobra.Command{
	Use:   "create-profile",
	Short: "Create a profile.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wellarchitected_createProfileCmd).Standalone()

	wellarchitected_createProfileCmd.Flags().String("client-request-token", "", "")
	wellarchitected_createProfileCmd.Flags().String("profile-description", "", "The profile description.")
	wellarchitected_createProfileCmd.Flags().String("profile-name", "", "Name of the profile.")
	wellarchitected_createProfileCmd.Flags().String("profile-questions", "", "The profile questions.")
	wellarchitected_createProfileCmd.Flags().String("tags", "", "The tags assigned to the profile.")
	wellarchitected_createProfileCmd.MarkFlagRequired("client-request-token")
	wellarchitected_createProfileCmd.MarkFlagRequired("profile-description")
	wellarchitected_createProfileCmd.MarkFlagRequired("profile-name")
	wellarchitected_createProfileCmd.MarkFlagRequired("profile-questions")
	wellarchitectedCmd.AddCommand(wellarchitected_createProfileCmd)
}
