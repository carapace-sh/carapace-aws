package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_createDataSourceCmd = &cobra.Command{
	Use:   "create-data-source",
	Short: "Creates a data source.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_createDataSourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(quicksight_createDataSourceCmd).Standalone()

		quicksight_createDataSourceCmd.Flags().String("aws-account-id", "", "The Amazon Web Services account ID.")
		quicksight_createDataSourceCmd.Flags().String("credentials", "", "The credentials Amazon Quick Sight that uses to connect to your underlying source.")
		quicksight_createDataSourceCmd.Flags().String("data-source-id", "", "An ID for the data source.")
		quicksight_createDataSourceCmd.Flags().String("data-source-parameters", "", "The parameters that Amazon Quick Sight uses to connect to your underlying source.")
		quicksight_createDataSourceCmd.Flags().String("folder-arns", "", "When you create the data source, Amazon Quick Sight adds the data source to these folders.")
		quicksight_createDataSourceCmd.Flags().String("name", "", "A display name for the data source.")
		quicksight_createDataSourceCmd.Flags().String("permissions", "", "A list of resource permissions on the data source.")
		quicksight_createDataSourceCmd.Flags().String("ssl-properties", "", "Secure Socket Layer (SSL) properties that apply when Amazon Quick Sight connects to your underlying source.")
		quicksight_createDataSourceCmd.Flags().String("tags", "", "Contains a map of the key-value pairs for the resource tag or tags assigned to the data source.")
		quicksight_createDataSourceCmd.Flags().String("type", "", "The type of the data source.")
		quicksight_createDataSourceCmd.Flags().String("vpc-connection-properties", "", "Use this parameter only when you want Amazon Quick Sight to use a VPC connection when connecting to your underlying source.")
		quicksight_createDataSourceCmd.MarkFlagRequired("aws-account-id")
		quicksight_createDataSourceCmd.MarkFlagRequired("data-source-id")
		quicksight_createDataSourceCmd.MarkFlagRequired("name")
		quicksight_createDataSourceCmd.MarkFlagRequired("type")
	})
	quicksightCmd.AddCommand(quicksight_createDataSourceCmd)
}
