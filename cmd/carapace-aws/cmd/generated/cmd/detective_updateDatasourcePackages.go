package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var detective_updateDatasourcePackagesCmd = &cobra.Command{
	Use:   "update-datasource-packages",
	Short: "Starts a data source package for the Detective behavior graph.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(detective_updateDatasourcePackagesCmd).Standalone()

	detective_updateDatasourcePackagesCmd.Flags().String("datasource-packages", "", "The data source package to start for the behavior graph.")
	detective_updateDatasourcePackagesCmd.Flags().String("graph-arn", "", "The ARN of the behavior graph.")
	detective_updateDatasourcePackagesCmd.MarkFlagRequired("datasource-packages")
	detective_updateDatasourcePackagesCmd.MarkFlagRequired("graph-arn")
	detectiveCmd.AddCommand(detective_updateDatasourcePackagesCmd)
}
