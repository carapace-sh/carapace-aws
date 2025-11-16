package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wellarchitected_deleteProfileShareCmd = &cobra.Command{
	Use:   "delete-profile-share",
	Short: "Delete a profile share.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wellarchitected_deleteProfileShareCmd).Standalone()

	wellarchitected_deleteProfileShareCmd.Flags().String("client-request-token", "", "")
	wellarchitected_deleteProfileShareCmd.Flags().String("profile-arn", "", "The profile ARN.")
	wellarchitected_deleteProfileShareCmd.Flags().String("share-id", "", "")
	wellarchitected_deleteProfileShareCmd.MarkFlagRequired("client-request-token")
	wellarchitected_deleteProfileShareCmd.MarkFlagRequired("profile-arn")
	wellarchitected_deleteProfileShareCmd.MarkFlagRequired("share-id")
	wellarchitectedCmd.AddCommand(wellarchitected_deleteProfileShareCmd)
}
