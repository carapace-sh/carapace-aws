package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wellarchitected_getProfileCmd = &cobra.Command{
	Use:   "get-profile",
	Short: "Get profile information.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wellarchitected_getProfileCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(wellarchitected_getProfileCmd).Standalone()

		wellarchitected_getProfileCmd.Flags().String("profile-arn", "", "The profile ARN.")
		wellarchitected_getProfileCmd.Flags().String("profile-version", "", "The profile version.")
		wellarchitected_getProfileCmd.MarkFlagRequired("profile-arn")
	})
	wellarchitectedCmd.AddCommand(wellarchitected_getProfileCmd)
}
