package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var imagebuilder_listInfrastructureConfigurationsCmd = &cobra.Command{
	Use:   "list-infrastructure-configurations",
	Short: "Returns a list of infrastructure configurations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(imagebuilder_listInfrastructureConfigurationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(imagebuilder_listInfrastructureConfigurationsCmd).Standalone()

		imagebuilder_listInfrastructureConfigurationsCmd.Flags().String("filters", "", "You can filter on `name` to streamline results.")
		imagebuilder_listInfrastructureConfigurationsCmd.Flags().String("max-results", "", "Specify the maximum number of items to return in a request.")
		imagebuilder_listInfrastructureConfigurationsCmd.Flags().String("next-token", "", "A token to specify where to start paginating.")
	})
	imagebuilderCmd.AddCommand(imagebuilder_listInfrastructureConfigurationsCmd)
}
