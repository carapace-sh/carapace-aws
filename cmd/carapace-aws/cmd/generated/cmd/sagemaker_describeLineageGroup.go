package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_describeLineageGroupCmd = &cobra.Command{
	Use:   "describe-lineage-group",
	Short: "Provides a list of properties for the requested lineage group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_describeLineageGroupCmd).Standalone()

	sagemaker_describeLineageGroupCmd.Flags().String("lineage-group-name", "", "The name of the lineage group.")
	sagemaker_describeLineageGroupCmd.MarkFlagRequired("lineage-group-name")
	sagemakerCmd.AddCommand(sagemaker_describeLineageGroupCmd)
}
