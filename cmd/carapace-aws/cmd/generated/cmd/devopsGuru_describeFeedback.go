package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var devopsGuru_describeFeedbackCmd = &cobra.Command{
	Use:   "describe-feedback",
	Short: "Returns the most recent feedback submitted in the current Amazon Web Services account and Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(devopsGuru_describeFeedbackCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(devopsGuru_describeFeedbackCmd).Standalone()

		devopsGuru_describeFeedbackCmd.Flags().String("insight-id", "", "The ID of the insight for which the feedback was provided.")
	})
	devopsGuruCmd.AddCommand(devopsGuru_describeFeedbackCmd)
}
