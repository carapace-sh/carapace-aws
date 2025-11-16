package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var support_describeSeverityLevelsCmd = &cobra.Command{
	Use:   "describe-severity-levels",
	Short: "Returns the list of severity levels that you can assign to a support case.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(support_describeSeverityLevelsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(support_describeSeverityLevelsCmd).Standalone()

		support_describeSeverityLevelsCmd.Flags().String("language", "", "The language in which Amazon Web Services Support handles the case.")
	})
	supportCmd.AddCommand(support_describeSeverityLevelsCmd)
}
