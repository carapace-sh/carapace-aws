package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_queryLineageCmd = &cobra.Command{
	Use:   "query-lineage",
	Short: "Use this action to inspect your lineage and discover relationships between entities.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_queryLineageCmd).Standalone()

	sagemaker_queryLineageCmd.Flags().String("direction", "", "Associations between lineage entities have a direction.")
	sagemaker_queryLineageCmd.Flags().String("filters", "", "A set of filtering parameters that allow you to specify which entities should be returned.")
	sagemaker_queryLineageCmd.Flags().Bool("include-edges", false, "Setting this value to `True` retrieves not only the entities of interest but also the [Associations](https://docs.aws.amazon.com/sagemaker/latest/dg/lineage-tracking-entities.html) and lineage entities on the path.")
	sagemaker_queryLineageCmd.Flags().String("max-depth", "", "The maximum depth in lineage relationships from the `StartArns` that are traversed.")
	sagemaker_queryLineageCmd.Flags().String("max-results", "", "Limits the number of vertices in the results.")
	sagemaker_queryLineageCmd.Flags().String("next-token", "", "Limits the number of vertices in the request.")
	sagemaker_queryLineageCmd.Flags().Bool("no-include-edges", false, "Setting this value to `True` retrieves not only the entities of interest but also the [Associations](https://docs.aws.amazon.com/sagemaker/latest/dg/lineage-tracking-entities.html) and lineage entities on the path.")
	sagemaker_queryLineageCmd.Flags().String("start-arns", "", "A list of resource Amazon Resource Name (ARN) that represent the starting point for your lineage query.")
	sagemaker_queryLineageCmd.Flag("no-include-edges").Hidden = true
	sagemakerCmd.AddCommand(sagemaker_queryLineageCmd)
}
