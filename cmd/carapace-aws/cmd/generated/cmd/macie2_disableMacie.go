package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var macie2_disableMacieCmd = &cobra.Command{
	Use:   "disable-macie",
	Short: "Disables Amazon Macie and deletes all settings and resources for a Macie account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(macie2_disableMacieCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(macie2_disableMacieCmd).Standalone()

	})
	macie2Cmd.AddCommand(macie2_disableMacieCmd)
}
