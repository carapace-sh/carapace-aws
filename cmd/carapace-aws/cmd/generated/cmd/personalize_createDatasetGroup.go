package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var personalize_createDatasetGroupCmd = &cobra.Command{
	Use:   "create-dataset-group",
	Short: "Creates an empty dataset group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(personalize_createDatasetGroupCmd).Standalone()

	personalize_createDatasetGroupCmd.Flags().String("domain", "", "The domain of the dataset group.")
	personalize_createDatasetGroupCmd.Flags().String("kms-key-arn", "", "The Amazon Resource Name (ARN) of a Key Management Service (KMS) key used to encrypt the datasets.")
	personalize_createDatasetGroupCmd.Flags().String("name", "", "The name for the new dataset group.")
	personalize_createDatasetGroupCmd.Flags().String("role-arn", "", "The ARN of the Identity and Access Management (IAM) role that has permissions to access the Key Management Service (KMS) key.")
	personalize_createDatasetGroupCmd.Flags().String("tags", "", "A list of [tags](https://docs.aws.amazon.com/personalize/latest/dg/tagging-resources.html) to apply to the dataset group.")
	personalize_createDatasetGroupCmd.MarkFlagRequired("name")
	personalizeCmd.AddCommand(personalize_createDatasetGroupCmd)
}
