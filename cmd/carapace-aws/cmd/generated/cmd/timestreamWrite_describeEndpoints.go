package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var timestreamWrite_describeEndpointsCmd = &cobra.Command{
	Use:   "describe-endpoints",
	Short: "Returns a list of available endpoints to make Timestream API calls against.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(timestreamWrite_describeEndpointsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(timestreamWrite_describeEndpointsCmd).Standalone()

	})
	timestreamWriteCmd.AddCommand(timestreamWrite_describeEndpointsCmd)
}
