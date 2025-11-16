package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_getConnectionsCmd = &cobra.Command{
	Use:   "get-connections",
	Short: "Retrieves a list of connection definitions from the Data Catalog.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_getConnectionsCmd).Standalone()

	glue_getConnectionsCmd.Flags().String("catalog-id", "", "The ID of the Data Catalog in which the connections reside.")
	glue_getConnectionsCmd.Flags().String("filter", "", "A filter that controls which connections are returned.")
	glue_getConnectionsCmd.Flags().Bool("hide-password", false, "Allows you to retrieve the connection metadata without returning the password.")
	glue_getConnectionsCmd.Flags().String("max-results", "", "The maximum number of connections to return in one response.")
	glue_getConnectionsCmd.Flags().String("next-token", "", "A continuation token, if this is a continuation call.")
	glue_getConnectionsCmd.Flags().Bool("no-hide-password", false, "Allows you to retrieve the connection metadata without returning the password.")
	glue_getConnectionsCmd.Flag("no-hide-password").Hidden = true
	glueCmd.AddCommand(glue_getConnectionsCmd)
}
