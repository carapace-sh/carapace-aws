package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediaconvert_getPolicyCmd = &cobra.Command{
	Use:   "get-policy",
	Short: "Retrieve the JSON for your policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediaconvert_getPolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mediaconvert_getPolicyCmd).Standalone()

	})
	mediaconvertCmd.AddCommand(mediaconvert_getPolicyCmd)
}
