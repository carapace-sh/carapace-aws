package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_updateDataSetCmd = &cobra.Command{
	Use:   "update-data-set",
	Short: "Updates a dataset.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_updateDataSetCmd).Standalone()

	quicksight_updateDataSetCmd.Flags().String("aws-account-id", "", "The Amazon Web Services account ID.")
	quicksight_updateDataSetCmd.Flags().String("column-groups", "", "Groupings of columns that work together in certain Amazon Quick Sight features.")
	quicksight_updateDataSetCmd.Flags().String("column-level-permission-rules", "", "A set of one or more definitions of a `ColumnLevelPermissionRule`.")
	quicksight_updateDataSetCmd.Flags().String("data-prep-configuration", "", "The data preparation configuration for the dataset.")
	quicksight_updateDataSetCmd.Flags().String("data-set-id", "", "The ID for the dataset that you want to update.")
	quicksight_updateDataSetCmd.Flags().String("data-set-usage-configuration", "", "")
	quicksight_updateDataSetCmd.Flags().String("dataset-parameters", "", "The parameter declarations of the dataset.")
	quicksight_updateDataSetCmd.Flags().String("field-folders", "", "The folder that contains fields and nested subfolders for your dataset.")
	quicksight_updateDataSetCmd.Flags().String("import-mode", "", "Indicates whether you want to import the data into SPICE.")
	quicksight_updateDataSetCmd.Flags().String("logical-table-map", "", "Configures the combination and transformation of the data from the physical tables.")
	quicksight_updateDataSetCmd.Flags().String("name", "", "The display name for the dataset.")
	quicksight_updateDataSetCmd.Flags().String("performance-configuration", "", "The configuration for the performance optimization of the dataset that contains a `UniqueKey` configuration.")
	quicksight_updateDataSetCmd.Flags().String("physical-table-map", "", "Declares the physical tables that are available in the underlying data sources.")
	quicksight_updateDataSetCmd.Flags().String("row-level-permission-data-set", "", "The row-level security configuration for the data you want to create.")
	quicksight_updateDataSetCmd.Flags().String("row-level-permission-tag-configuration", "", "The configuration of tags on a dataset to set row-level security.")
	quicksight_updateDataSetCmd.Flags().String("semantic-model-configuration", "", "The semantic model configuration for the dataset.")
	quicksight_updateDataSetCmd.MarkFlagRequired("aws-account-id")
	quicksight_updateDataSetCmd.MarkFlagRequired("data-set-id")
	quicksight_updateDataSetCmd.MarkFlagRequired("import-mode")
	quicksight_updateDataSetCmd.MarkFlagRequired("name")
	quicksight_updateDataSetCmd.MarkFlagRequired("physical-table-map")
	quicksightCmd.AddCommand(quicksight_updateDataSetCmd)
}
