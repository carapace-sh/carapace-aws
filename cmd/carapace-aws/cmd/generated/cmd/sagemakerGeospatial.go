package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemakerGeospatialCmd = &cobra.Command{
	Use:   "sagemaker-geospatial",
	Short: "Provides APIs for creating and managing SageMaker geospatial resources.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemakerGeospatialCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemakerGeospatialCmd).Standalone()

	})
	rootCmd.AddCommand(sagemakerGeospatialCmd)
}
