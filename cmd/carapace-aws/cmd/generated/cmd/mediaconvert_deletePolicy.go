package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediaconvert_deletePolicyCmd = &cobra.Command{
	Use:   "delete-policy",
	Short: "Permanently delete a policy that you created.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediaconvert_deletePolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mediaconvert_deletePolicyCmd).Standalone()

	})
	mediaconvertCmd.AddCommand(mediaconvert_deletePolicyCmd)
}
