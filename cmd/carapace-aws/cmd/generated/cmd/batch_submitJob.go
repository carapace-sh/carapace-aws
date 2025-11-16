package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var batch_submitJobCmd = &cobra.Command{
	Use:   "submit-job",
	Short: "Submits an Batch job from a job definition.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(batch_submitJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(batch_submitJobCmd).Standalone()

		batch_submitJobCmd.Flags().String("array-properties", "", "The array properties for the submitted job, such as the size of the array.")
		batch_submitJobCmd.Flags().String("consumable-resource-properties-override", "", "An object that contains overrides for the consumable resources of a job.")
		batch_submitJobCmd.Flags().String("container-overrides", "", "An object with properties that override the defaults for the job definition that specify the name of a container in the specified job definition and the overrides it should receive.")
		batch_submitJobCmd.Flags().String("depends-on", "", "A list of dependencies for the job.")
		batch_submitJobCmd.Flags().String("ecs-properties-override", "", "An object, with properties that override defaults for the job definition, can only be specified for jobs that are run on Amazon ECS resources.")
		batch_submitJobCmd.Flags().String("eks-properties-override", "", "An object, with properties that override defaults for the job definition, can only be specified for jobs that are run on Amazon EKS resources.")
		batch_submitJobCmd.Flags().String("job-definition", "", "The job definition used by this job.")
		batch_submitJobCmd.Flags().String("job-name", "", "The name of the job.")
		batch_submitJobCmd.Flags().String("job-queue", "", "The job queue where the job is submitted.")
		batch_submitJobCmd.Flags().Bool("no-propagate-tags", false, "Specifies whether to propagate the tags from the job or job definition to the corresponding Amazon ECS task.")
		batch_submitJobCmd.Flags().String("node-overrides", "", "A list of node overrides in JSON format that specify the node range to target and the container overrides for that node range.")
		batch_submitJobCmd.Flags().String("parameters", "", "Additional parameters passed to the job that replace parameter substitution placeholders that are set in the job definition.")
		batch_submitJobCmd.Flags().Bool("propagate-tags", false, "Specifies whether to propagate the tags from the job or job definition to the corresponding Amazon ECS task.")
		batch_submitJobCmd.Flags().String("retry-strategy", "", "The retry strategy to use for failed jobs from this [SubmitJob]() operation.")
		batch_submitJobCmd.Flags().String("scheduling-priority-override", "", "The scheduling priority for the job.")
		batch_submitJobCmd.Flags().String("share-identifier", "", "The share identifier for the job.")
		batch_submitJobCmd.Flags().String("tags", "", "The tags that you apply to the job request to help you categorize and organize your resources.")
		batch_submitJobCmd.Flags().String("timeout", "", "The timeout configuration for this [SubmitJob]() operation.")
		batch_submitJobCmd.MarkFlagRequired("job-definition")
		batch_submitJobCmd.MarkFlagRequired("job-name")
		batch_submitJobCmd.MarkFlagRequired("job-queue")
		batch_submitJobCmd.Flag("no-propagate-tags").Hidden = true
	})
	batchCmd.AddCommand(batch_submitJobCmd)
}
