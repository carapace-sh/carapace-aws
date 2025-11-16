package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iottwinmaker_listComponentTypesCmd = &cobra.Command{
	Use:   "list-component-types",
	Short: "Lists all component types in a workspace.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iottwinmaker_listComponentTypesCmd).Standalone()

	iottwinmaker_listComponentTypesCmd.Flags().String("filters", "", "A list of objects that filter the request.")
	iottwinmaker_listComponentTypesCmd.Flags().String("max-results", "", "The maximum number of results to return at one time.")
	iottwinmaker_listComponentTypesCmd.Flags().String("next-token", "", "The string that specifies the next page of results.")
	iottwinmaker_listComponentTypesCmd.Flags().String("workspace-id", "", "The ID of the workspace.")
	iottwinmaker_listComponentTypesCmd.MarkFlagRequired("workspace-id")
	iottwinmakerCmd.AddCommand(iottwinmaker_listComponentTypesCmd)
}
