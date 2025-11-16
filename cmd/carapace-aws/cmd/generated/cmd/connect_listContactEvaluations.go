package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_listContactEvaluationsCmd = &cobra.Command{
	Use:   "list-contact-evaluations",
	Short: "Lists contact evaluations in the specified Amazon Connect instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_listContactEvaluationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_listContactEvaluationsCmd).Standalone()

		connect_listContactEvaluationsCmd.Flags().String("contact-id", "", "The identifier of the contact in this instance of Amazon Connect.")
		connect_listContactEvaluationsCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_listContactEvaluationsCmd.Flags().String("next-token", "", "The token for the next set of results.")
		connect_listContactEvaluationsCmd.MarkFlagRequired("contact-id")
		connect_listContactEvaluationsCmd.MarkFlagRequired("instance-id")
	})
	connectCmd.AddCommand(connect_listContactEvaluationsCmd)
}
