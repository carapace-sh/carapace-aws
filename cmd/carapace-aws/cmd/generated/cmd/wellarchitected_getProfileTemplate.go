package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wellarchitected_getProfileTemplateCmd = &cobra.Command{
	Use:   "get-profile-template",
	Short: "Get profile template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wellarchitected_getProfileTemplateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(wellarchitected_getProfileTemplateCmd).Standalone()

	})
	wellarchitectedCmd.AddCommand(wellarchitected_getProfileTemplateCmd)
}
