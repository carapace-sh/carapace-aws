package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotfleetwise_getLoggingOptionsCmd = &cobra.Command{
	Use:   "get-logging-options",
	Short: "Retrieves the logging options.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotfleetwise_getLoggingOptionsCmd).Standalone()

	iotfleetwiseCmd.AddCommand(iotfleetwise_getLoggingOptionsCmd)
}
