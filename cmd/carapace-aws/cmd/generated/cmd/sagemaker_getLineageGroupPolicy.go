package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_getLineageGroupPolicyCmd = &cobra.Command{
	Use:   "get-lineage-group-policy",
	Short: "The resource policy for the lineage group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_getLineageGroupPolicyCmd).Standalone()

	sagemaker_getLineageGroupPolicyCmd.Flags().String("lineage-group-name", "", "The name or Amazon Resource Name (ARN) of the lineage group.")
	sagemaker_getLineageGroupPolicyCmd.MarkFlagRequired("lineage-group-name")
	sagemakerCmd.AddCommand(sagemaker_getLineageGroupPolicyCmd)
}
