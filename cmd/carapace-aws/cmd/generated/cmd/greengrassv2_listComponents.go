package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var greengrassv2_listComponentsCmd = &cobra.Command{
	Use:   "list-components",
	Short: "Retrieves a paginated list of component summaries.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(greengrassv2_listComponentsCmd).Standalone()

	greengrassv2_listComponentsCmd.Flags().String("max-results", "", "The maximum number of results to be returned per paginated request.")
	greengrassv2_listComponentsCmd.Flags().String("next-token", "", "The token to be used for the next set of paginated results.")
	greengrassv2_listComponentsCmd.Flags().String("scope", "", "The scope of the components to list.")
	greengrassv2Cmd.AddCommand(greengrassv2_listComponentsCmd)
}
