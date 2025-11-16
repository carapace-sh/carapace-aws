package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_batchGetDevEndpointsCmd = &cobra.Command{
	Use:   "batch-get-dev-endpoints",
	Short: "Returns a list of resource metadata for a given list of development endpoint names.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_batchGetDevEndpointsCmd).Standalone()

	glue_batchGetDevEndpointsCmd.Flags().String("dev-endpoint-names", "", "The list of `DevEndpoint` names, which might be the names returned from the `ListDevEndpoint` operation.")
	glue_batchGetDevEndpointsCmd.MarkFlagRequired("dev-endpoint-names")
	glueCmd.AddCommand(glue_batchGetDevEndpointsCmd)
}
