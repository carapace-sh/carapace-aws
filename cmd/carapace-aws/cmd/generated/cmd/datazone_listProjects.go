package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_listProjectsCmd = &cobra.Command{
	Use:   "list-projects",
	Short: "Lists Amazon DataZone projects.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_listProjectsCmd).Standalone()

	datazone_listProjectsCmd.Flags().String("domain-identifier", "", "The identifier of the Amazon DataZone domain.")
	datazone_listProjectsCmd.Flags().String("group-identifier", "", "The identifier of a group.")
	datazone_listProjectsCmd.Flags().String("max-results", "", "The maximum number of projects to return in a single call to `ListProjects`.")
	datazone_listProjectsCmd.Flags().String("name", "", "The name of the project.")
	datazone_listProjectsCmd.Flags().String("next-token", "", "When the number of projects is greater than the default value for the `MaxResults` parameter, or if you explicitly specify a value for `MaxResults` that is less than the number of projects, the response includes a pagination token named `NextToken`.")
	datazone_listProjectsCmd.Flags().String("user-identifier", "", "The identifier of the Amazon DataZone user.")
	datazone_listProjectsCmd.MarkFlagRequired("domain-identifier")
	datazoneCmd.AddCommand(datazone_listProjectsCmd)
}
