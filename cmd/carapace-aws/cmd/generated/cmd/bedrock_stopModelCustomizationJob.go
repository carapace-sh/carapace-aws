package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrock_stopModelCustomizationJobCmd = &cobra.Command{
	Use:   "stop-model-customization-job",
	Short: "Stops an active model customization job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrock_stopModelCustomizationJobCmd).Standalone()

	bedrock_stopModelCustomizationJobCmd.Flags().String("job-identifier", "", "Job identifier of the job to stop.")
	bedrock_stopModelCustomizationJobCmd.MarkFlagRequired("job-identifier")
	bedrockCmd.AddCommand(bedrock_stopModelCustomizationJobCmd)
}
