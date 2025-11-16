package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_startContactEvaluationCmd = &cobra.Command{
	Use:   "start-contact-evaluation",
	Short: "Starts an empty evaluation in the specified Amazon Connect instance, using the given evaluation form for the particular contact.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_startContactEvaluationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_startContactEvaluationCmd).Standalone()

		connect_startContactEvaluationCmd.Flags().String("auto-evaluation-configuration", "", "Whether automated evaluations are enabled.")
		connect_startContactEvaluationCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		connect_startContactEvaluationCmd.Flags().String("contact-id", "", "The identifier of the contact in this instance of Amazon Connect.")
		connect_startContactEvaluationCmd.Flags().String("evaluation-form-id", "", "The unique identifier for the evaluation form.")
		connect_startContactEvaluationCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_startContactEvaluationCmd.Flags().String("tags", "", "The tags used to organize, track, or control access for this resource.")
		connect_startContactEvaluationCmd.MarkFlagRequired("contact-id")
		connect_startContactEvaluationCmd.MarkFlagRequired("evaluation-form-id")
		connect_startContactEvaluationCmd.MarkFlagRequired("instance-id")
	})
	connectCmd.AddCommand(connect_startContactEvaluationCmd)
}
