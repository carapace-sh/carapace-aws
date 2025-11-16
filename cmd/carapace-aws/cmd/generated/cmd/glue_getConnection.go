package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_getConnectionCmd = &cobra.Command{
	Use:   "get-connection",
	Short: "Retrieves a connection definition from the Data Catalog.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_getConnectionCmd).Standalone()

	glue_getConnectionCmd.Flags().String("apply-override-for-compute-environment", "", "For connections that may be used in multiple services, specifies returning properties for the specified compute environment.")
	glue_getConnectionCmd.Flags().String("catalog-id", "", "The ID of the Data Catalog in which the connection resides.")
	glue_getConnectionCmd.Flags().Bool("hide-password", false, "Allows you to retrieve the connection metadata without returning the password.")
	glue_getConnectionCmd.Flags().String("name", "", "The name of the connection definition to retrieve.")
	glue_getConnectionCmd.Flags().Bool("no-hide-password", false, "Allows you to retrieve the connection metadata without returning the password.")
	glue_getConnectionCmd.MarkFlagRequired("name")
	glue_getConnectionCmd.Flag("no-hide-password").Hidden = true
	glueCmd.AddCommand(glue_getConnectionCmd)
}
