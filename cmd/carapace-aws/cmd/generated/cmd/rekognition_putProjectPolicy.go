package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rekognition_putProjectPolicyCmd = &cobra.Command{
	Use:   "put-project-policy",
	Short: "This operation applies only to Amazon Rekognition Custom Labels.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rekognition_putProjectPolicyCmd).Standalone()

	rekognition_putProjectPolicyCmd.Flags().String("policy-document", "", "A resource policy to add to the model.")
	rekognition_putProjectPolicyCmd.Flags().String("policy-name", "", "A name for the policy.")
	rekognition_putProjectPolicyCmd.Flags().String("policy-revision-id", "", "The revision ID for the Project Policy.")
	rekognition_putProjectPolicyCmd.Flags().String("project-arn", "", "The Amazon Resource Name (ARN) of the project that the project policy is attached to.")
	rekognition_putProjectPolicyCmd.MarkFlagRequired("policy-document")
	rekognition_putProjectPolicyCmd.MarkFlagRequired("policy-name")
	rekognition_putProjectPolicyCmd.MarkFlagRequired("project-arn")
	rekognitionCmd.AddCommand(rekognition_putProjectPolicyCmd)
}
