package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_listConnectionsCmd = &cobra.Command{
	Use:   "list-connections",
	Short: "Lists connections.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_listConnectionsCmd).Standalone()

	datazone_listConnectionsCmd.Flags().String("domain-identifier", "", "The ID of the domain where you want to list connections.")
	datazone_listConnectionsCmd.Flags().String("environment-identifier", "", "The ID of the environment where you want to list connections.")
	datazone_listConnectionsCmd.Flags().String("max-results", "", "The maximum number of connections to return in a single call to ListConnections.")
	datazone_listConnectionsCmd.Flags().String("name", "", "The name of the connection.")
	datazone_listConnectionsCmd.Flags().String("next-token", "", "When the number of connections is greater than the default value for the MaxResults parameter, or if you explicitly specify a value for MaxResults that is less than the number of connections, the response includes a pagination token named NextToken.")
	datazone_listConnectionsCmd.Flags().String("project-identifier", "", "The ID of the project where you want to list connections.")
	datazone_listConnectionsCmd.Flags().String("scope", "", "The scope of the connection.")
	datazone_listConnectionsCmd.Flags().String("sort-by", "", "Specifies how you want to sort the listed connections.")
	datazone_listConnectionsCmd.Flags().String("sort-order", "", "Specifies the sort order for the listed connections.")
	datazone_listConnectionsCmd.Flags().String("type", "", "The type of connection.")
	datazone_listConnectionsCmd.MarkFlagRequired("domain-identifier")
	datazoneCmd.AddCommand(datazone_listConnectionsCmd)
}
