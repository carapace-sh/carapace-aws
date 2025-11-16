package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var imagebuilder_listDistributionConfigurationsCmd = &cobra.Command{
	Use:   "list-distribution-configurations",
	Short: "Returns a list of distribution configurations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(imagebuilder_listDistributionConfigurationsCmd).Standalone()

	imagebuilder_listDistributionConfigurationsCmd.Flags().String("filters", "", "You can filter on `name` to streamline results.")
	imagebuilder_listDistributionConfigurationsCmd.Flags().String("max-results", "", "Specify the maximum number of items to return in a request.")
	imagebuilder_listDistributionConfigurationsCmd.Flags().String("next-token", "", "A token to specify where to start paginating.")
	imagebuilderCmd.AddCommand(imagebuilder_listDistributionConfigurationsCmd)
}
