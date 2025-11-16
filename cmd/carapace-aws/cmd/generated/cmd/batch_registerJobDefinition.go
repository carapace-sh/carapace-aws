package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var batch_registerJobDefinitionCmd = &cobra.Command{
	Use:   "register-job-definition",
	Short: "Registers an Batch job definition.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(batch_registerJobDefinitionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(batch_registerJobDefinitionCmd).Standalone()

		batch_registerJobDefinitionCmd.Flags().String("consumable-resource-properties", "", "Contains a list of consumable resources required by the job.")
		batch_registerJobDefinitionCmd.Flags().String("container-properties", "", "An object with properties specific to Amazon ECS-based single-node container-based jobs.")
		batch_registerJobDefinitionCmd.Flags().String("ecs-properties", "", "An object with properties that are specific to Amazon ECS-based jobs.")
		batch_registerJobDefinitionCmd.Flags().String("eks-properties", "", "An object with properties that are specific to Amazon EKS-based jobs.")
		batch_registerJobDefinitionCmd.Flags().String("job-definition-name", "", "The name of the job definition to register.")
		batch_registerJobDefinitionCmd.Flags().Bool("no-propagate-tags", false, "Specifies whether to propagate the tags from the job or job definition to the corresponding Amazon ECS task.")
		batch_registerJobDefinitionCmd.Flags().String("node-properties", "", "An object with properties specific to multi-node parallel jobs.")
		batch_registerJobDefinitionCmd.Flags().String("parameters", "", "Default parameter substitution placeholders to set in the job definition.")
		batch_registerJobDefinitionCmd.Flags().String("platform-capabilities", "", "The platform capabilities required by the job definition.")
		batch_registerJobDefinitionCmd.Flags().Bool("propagate-tags", false, "Specifies whether to propagate the tags from the job or job definition to the corresponding Amazon ECS task.")
		batch_registerJobDefinitionCmd.Flags().String("retry-strategy", "", "The retry strategy to use for failed jobs that are submitted with this job definition.")
		batch_registerJobDefinitionCmd.Flags().String("scheduling-priority", "", "The scheduling priority for jobs that are submitted with this job definition.")
		batch_registerJobDefinitionCmd.Flags().String("tags", "", "The tags that you apply to the job definition to help you categorize and organize your resources.")
		batch_registerJobDefinitionCmd.Flags().String("timeout", "", "The timeout configuration for jobs that are submitted with this job definition, after which Batch terminates your jobs if they have not finished.")
		batch_registerJobDefinitionCmd.Flags().String("type", "", "The type of job definition.")
		batch_registerJobDefinitionCmd.MarkFlagRequired("job-definition-name")
		batch_registerJobDefinitionCmd.Flag("no-propagate-tags").Hidden = true
		batch_registerJobDefinitionCmd.MarkFlagRequired("type")
	})
	batchCmd.AddCommand(batch_registerJobDefinitionCmd)
}
