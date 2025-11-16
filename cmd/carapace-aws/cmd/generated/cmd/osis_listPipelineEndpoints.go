package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var osis_listPipelineEndpointsCmd = &cobra.Command{
	Use:   "list-pipeline-endpoints",
	Short: "Lists all pipeline endpoints in your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(osis_listPipelineEndpointsCmd).Standalone()

	osis_listPipelineEndpointsCmd.Flags().String("max-results", "", "The maximum number of pipeline endpoints to return in the response.")
	osis_listPipelineEndpointsCmd.Flags().String("next-token", "", "If your initial `ListPipelineEndpoints` operation returns a `NextToken`, you can include the returned `NextToken` in subsequent `ListPipelineEndpoints` operations, which returns results in the next page.")
	osisCmd.AddCommand(osis_listPipelineEndpointsCmd)
}
