package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var timestreamQuery_describeEndpointsCmd = &cobra.Command{
	Use:   "describe-endpoints",
	Short: "DescribeEndpoints returns a list of available endpoints to make Timestream API calls against.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(timestreamQuery_describeEndpointsCmd).Standalone()

	timestreamQueryCmd.AddCommand(timestreamQuery_describeEndpointsCmd)
}
