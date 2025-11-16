package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrock_getModelCustomizationJobCmd = &cobra.Command{
	Use:   "get-model-customization-job",
	Short: "Retrieves the properties associated with a model-customization job, including the status of the job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrock_getModelCustomizationJobCmd).Standalone()

	bedrock_getModelCustomizationJobCmd.Flags().String("job-identifier", "", "Identifier for the customization job.")
	bedrock_getModelCustomizationJobCmd.MarkFlagRequired("job-identifier")
	bedrockCmd.AddCommand(bedrock_getModelCustomizationJobCmd)
}
