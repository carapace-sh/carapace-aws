package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var macie2_getRevealConfigurationCmd = &cobra.Command{
	Use:   "get-reveal-configuration",
	Short: "Retrieves the status and configuration settings for retrieving occurrences of sensitive data reported by findings.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(macie2_getRevealConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(macie2_getRevealConfigurationCmd).Standalone()

	})
	macie2Cmd.AddCommand(macie2_getRevealConfigurationCmd)
}
