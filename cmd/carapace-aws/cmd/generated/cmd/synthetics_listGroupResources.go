package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var synthetics_listGroupResourcesCmd = &cobra.Command{
	Use:   "list-group-resources",
	Short: "This operation returns a list of the ARNs of the canaries that are associated with the specified group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(synthetics_listGroupResourcesCmd).Standalone()

	synthetics_listGroupResourcesCmd.Flags().String("group-identifier", "", "Specifies the group to return information for.")
	synthetics_listGroupResourcesCmd.Flags().String("max-results", "", "Specify this parameter to limit how many canary ARNs are returned each time you use the `ListGroupResources` operation.")
	synthetics_listGroupResourcesCmd.Flags().String("next-token", "", "A token that indicates that there is more data available.")
	synthetics_listGroupResourcesCmd.MarkFlagRequired("group-identifier")
	syntheticsCmd.AddCommand(synthetics_listGroupResourcesCmd)
}
