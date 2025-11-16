package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wellarchitected_createProfileShareCmd = &cobra.Command{
	Use:   "create-profile-share",
	Short: "Create a profile share.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wellarchitected_createProfileShareCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(wellarchitected_createProfileShareCmd).Standalone()

		wellarchitected_createProfileShareCmd.Flags().String("client-request-token", "", "")
		wellarchitected_createProfileShareCmd.Flags().String("profile-arn", "", "The profile ARN.")
		wellarchitected_createProfileShareCmd.Flags().String("shared-with", "", "")
		wellarchitected_createProfileShareCmd.MarkFlagRequired("client-request-token")
		wellarchitected_createProfileShareCmd.MarkFlagRequired("profile-arn")
		wellarchitected_createProfileShareCmd.MarkFlagRequired("shared-with")
	})
	wellarchitectedCmd.AddCommand(wellarchitected_createProfileShareCmd)
}
