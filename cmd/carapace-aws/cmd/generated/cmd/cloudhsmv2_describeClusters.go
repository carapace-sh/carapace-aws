package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudhsmv2_describeClustersCmd = &cobra.Command{
	Use:   "describe-clusters",
	Short: "Gets information about CloudHSM clusters.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudhsmv2_describeClustersCmd).Standalone()

	cloudhsmv2_describeClustersCmd.Flags().String("filters", "", "One or more filters to limit the items returned in the response.")
	cloudhsmv2_describeClustersCmd.Flags().String("max-results", "", "The maximum number of clusters to return in the response.")
	cloudhsmv2_describeClustersCmd.Flags().String("next-token", "", "The `NextToken` value that you received in the previous response.")
	cloudhsmv2Cmd.AddCommand(cloudhsmv2_describeClustersCmd)
}
