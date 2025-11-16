package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_updateDataSourceCmd = &cobra.Command{
	Use:   "update-data-source",
	Short: "Updates a data source.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_updateDataSourceCmd).Standalone()

	quicksight_updateDataSourceCmd.Flags().String("aws-account-id", "", "The Amazon Web Services account ID.")
	quicksight_updateDataSourceCmd.Flags().String("credentials", "", "The credentials that Amazon Quick Sight that uses to connect to your underlying source.")
	quicksight_updateDataSourceCmd.Flags().String("data-source-id", "", "The ID of the data source.")
	quicksight_updateDataSourceCmd.Flags().String("data-source-parameters", "", "The parameters that Amazon Quick Sight uses to connect to your underlying source.")
	quicksight_updateDataSourceCmd.Flags().String("name", "", "A display name for the data source.")
	quicksight_updateDataSourceCmd.Flags().String("ssl-properties", "", "Secure Socket Layer (SSL) properties that apply when Amazon Quick Sight connects to your underlying source.")
	quicksight_updateDataSourceCmd.Flags().String("vpc-connection-properties", "", "Use this parameter only when you want Amazon Quick Sight to use a VPC connection when connecting to your underlying source.")
	quicksight_updateDataSourceCmd.MarkFlagRequired("aws-account-id")
	quicksight_updateDataSourceCmd.MarkFlagRequired("data-source-id")
	quicksight_updateDataSourceCmd.MarkFlagRequired("name")
	quicksightCmd.AddCommand(quicksight_updateDataSourceCmd)
}
