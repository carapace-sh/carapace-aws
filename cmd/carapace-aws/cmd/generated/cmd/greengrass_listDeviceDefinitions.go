package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var greengrass_listDeviceDefinitionsCmd = &cobra.Command{
	Use:   "list-device-definitions",
	Short: "Retrieves a list of device definitions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(greengrass_listDeviceDefinitionsCmd).Standalone()

	greengrass_listDeviceDefinitionsCmd.Flags().String("max-results", "", "The maximum number of results to be returned per request.")
	greengrass_listDeviceDefinitionsCmd.Flags().String("next-token", "", "The token for the next set of results, or ''null'' if there are no additional results.")
	greengrassCmd.AddCommand(greengrass_listDeviceDefinitionsCmd)
}
