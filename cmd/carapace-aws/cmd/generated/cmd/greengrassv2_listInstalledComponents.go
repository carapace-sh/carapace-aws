package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var greengrassv2_listInstalledComponentsCmd = &cobra.Command{
	Use:   "list-installed-components",
	Short: "Retrieves a paginated list of the components that a Greengrass core device runs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(greengrassv2_listInstalledComponentsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(greengrassv2_listInstalledComponentsCmd).Standalone()

		greengrassv2_listInstalledComponentsCmd.Flags().String("core-device-thing-name", "", "The name of the core device.")
		greengrassv2_listInstalledComponentsCmd.Flags().String("max-results", "", "The maximum number of results to be returned per paginated request.")
		greengrassv2_listInstalledComponentsCmd.Flags().String("next-token", "", "The token to be used for the next set of paginated results.")
		greengrassv2_listInstalledComponentsCmd.Flags().String("topology-filter", "", "The filter for the list of components.")
		greengrassv2_listInstalledComponentsCmd.MarkFlagRequired("core-device-thing-name")
	})
	greengrassv2Cmd.AddCommand(greengrassv2_listInstalledComponentsCmd)
}
