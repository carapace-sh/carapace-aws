package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_createDataSetCmd = &cobra.Command{
	Use:   "create-data-set",
	Short: "Creates a dataset.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_createDataSetCmd).Standalone()

	quicksight_createDataSetCmd.Flags().String("aws-account-id", "", "The Amazon Web Services account ID.")
	quicksight_createDataSetCmd.Flags().String("column-groups", "", "Groupings of columns that work together in certain Amazon Quick Sight features.")
	quicksight_createDataSetCmd.Flags().String("column-level-permission-rules", "", "A set of one or more definitions of a `ColumnLevelPermissionRule`.")
	quicksight_createDataSetCmd.Flags().String("data-prep-configuration", "", "The data preparation configuration for the dataset.")
	quicksight_createDataSetCmd.Flags().String("data-set-id", "", "An ID for the dataset that you want to create.")
	quicksight_createDataSetCmd.Flags().String("data-set-usage-configuration", "", "")
	quicksight_createDataSetCmd.Flags().String("dataset-parameters", "", "The parameter declarations of the dataset.")
	quicksight_createDataSetCmd.Flags().String("field-folders", "", "The folder that contains fields and nested subfolders for your dataset.")
	quicksight_createDataSetCmd.Flags().String("folder-arns", "", "When you create the dataset, Amazon Quick Sight adds the dataset to these folders.")
	quicksight_createDataSetCmd.Flags().String("import-mode", "", "Indicates whether you want to import the data into SPICE.")
	quicksight_createDataSetCmd.Flags().String("logical-table-map", "", "Configures the combination and transformation of the data from the physical tables.")
	quicksight_createDataSetCmd.Flags().String("name", "", "The display name for the dataset.")
	quicksight_createDataSetCmd.Flags().String("performance-configuration", "", "The configuration for the performance optimization of the dataset that contains a `UniqueKey` configuration.")
	quicksight_createDataSetCmd.Flags().String("permissions", "", "A list of resource permissions on the dataset.")
	quicksight_createDataSetCmd.Flags().String("physical-table-map", "", "Declares the physical tables that are available in the underlying data sources.")
	quicksight_createDataSetCmd.Flags().String("row-level-permission-data-set", "", "The row-level security configuration for the data that you want to create.")
	quicksight_createDataSetCmd.Flags().String("row-level-permission-tag-configuration", "", "The configuration of tags on a dataset to set row-level security.")
	quicksight_createDataSetCmd.Flags().String("semantic-model-configuration", "", "The semantic model configuration for the dataset.")
	quicksight_createDataSetCmd.Flags().String("tags", "", "Contains a map of the key-value pairs for the resource tag or tags assigned to the dataset.")
	quicksight_createDataSetCmd.Flags().String("use-as", "", "The usage of the dataset.")
	quicksight_createDataSetCmd.MarkFlagRequired("aws-account-id")
	quicksight_createDataSetCmd.MarkFlagRequired("data-set-id")
	quicksight_createDataSetCmd.MarkFlagRequired("import-mode")
	quicksight_createDataSetCmd.MarkFlagRequired("name")
	quicksight_createDataSetCmd.MarkFlagRequired("physical-table-map")
	quicksightCmd.AddCommand(quicksight_createDataSetCmd)
}
