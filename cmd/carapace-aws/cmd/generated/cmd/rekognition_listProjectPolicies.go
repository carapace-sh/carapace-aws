package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rekognition_listProjectPoliciesCmd = &cobra.Command{
	Use:   "list-project-policies",
	Short: "This operation applies only to Amazon Rekognition Custom Labels.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rekognition_listProjectPoliciesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rekognition_listProjectPoliciesCmd).Standalone()

		rekognition_listProjectPoliciesCmd.Flags().String("max-results", "", "The maximum number of results to return per paginated call.")
		rekognition_listProjectPoliciesCmd.Flags().String("next-token", "", "If the previous response was incomplete (because there is more results to retrieve), Amazon Rekognition Custom Labels returns a pagination token in the response.")
		rekognition_listProjectPoliciesCmd.Flags().String("project-arn", "", "The ARN of the project for which you want to list the project policies.")
		rekognition_listProjectPoliciesCmd.MarkFlagRequired("project-arn")
	})
	rekognitionCmd.AddCommand(rekognition_listProjectPoliciesCmd)
}
