package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wellarchitected_deleteProfileCmd = &cobra.Command{
	Use:   "delete-profile",
	Short: "Delete a profile.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wellarchitected_deleteProfileCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(wellarchitected_deleteProfileCmd).Standalone()

		wellarchitected_deleteProfileCmd.Flags().String("client-request-token", "", "")
		wellarchitected_deleteProfileCmd.Flags().String("profile-arn", "", "The profile ARN.")
		wellarchitected_deleteProfileCmd.MarkFlagRequired("client-request-token")
		wellarchitected_deleteProfileCmd.MarkFlagRequired("profile-arn")
	})
	wellarchitectedCmd.AddCommand(wellarchitected_deleteProfileCmd)
}
