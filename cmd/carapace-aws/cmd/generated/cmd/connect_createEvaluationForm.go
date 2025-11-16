package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_createEvaluationFormCmd = &cobra.Command{
	Use:   "create-evaluation-form",
	Short: "Creates an evaluation form in the specified Amazon Connect instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_createEvaluationFormCmd).Standalone()

	connect_createEvaluationFormCmd.Flags().String("auto-evaluation-configuration", "", "Configuration information about automated evaluations.")
	connect_createEvaluationFormCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	connect_createEvaluationFormCmd.Flags().String("description", "", "The description of the evaluation form.")
	connect_createEvaluationFormCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_createEvaluationFormCmd.Flags().String("items", "", "Items that are part of the evaluation form.")
	connect_createEvaluationFormCmd.Flags().String("scoring-strategy", "", "A scoring strategy of the evaluation form.")
	connect_createEvaluationFormCmd.Flags().String("tags", "", "The tags used to organize, track, or control access for this resource.")
	connect_createEvaluationFormCmd.Flags().String("title", "", "A title of the evaluation form.")
	connect_createEvaluationFormCmd.MarkFlagRequired("instance-id")
	connect_createEvaluationFormCmd.MarkFlagRequired("items")
	connect_createEvaluationFormCmd.MarkFlagRequired("title")
	connectCmd.AddCommand(connect_createEvaluationFormCmd)
}
