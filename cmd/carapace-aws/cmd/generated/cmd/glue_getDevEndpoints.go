package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_getDevEndpointsCmd = &cobra.Command{
	Use:   "get-dev-endpoints",
	Short: "Retrieves all the development endpoints in this Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_getDevEndpointsCmd).Standalone()

	glue_getDevEndpointsCmd.Flags().String("max-results", "", "The maximum size of information to return.")
	glue_getDevEndpointsCmd.Flags().String("next-token", "", "A continuation token, if this is a continuation call.")
	glueCmd.AddCommand(glue_getDevEndpointsCmd)
}
