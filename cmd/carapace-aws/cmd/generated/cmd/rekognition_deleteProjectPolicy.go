package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rekognition_deleteProjectPolicyCmd = &cobra.Command{
	Use:   "delete-project-policy",
	Short: "This operation applies only to Amazon Rekognition Custom Labels.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rekognition_deleteProjectPolicyCmd).Standalone()

	rekognition_deleteProjectPolicyCmd.Flags().String("policy-name", "", "The name of the policy that you want to delete.")
	rekognition_deleteProjectPolicyCmd.Flags().String("policy-revision-id", "", "The ID of the project policy revision that you want to delete.")
	rekognition_deleteProjectPolicyCmd.Flags().String("project-arn", "", "The Amazon Resource Name (ARN) of the project that the project policy you want to delete is attached to.")
	rekognition_deleteProjectPolicyCmd.MarkFlagRequired("policy-name")
	rekognition_deleteProjectPolicyCmd.MarkFlagRequired("project-arn")
	rekognitionCmd.AddCommand(rekognition_deleteProjectPolicyCmd)
}
