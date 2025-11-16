package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotevents_listInputsCmd = &cobra.Command{
	Use:   "list-inputs",
	Short: "Lists the inputs you have created.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotevents_listInputsCmd).Standalone()

	iotevents_listInputsCmd.Flags().String("max-results", "", "The maximum number of results to be returned per request.")
	iotevents_listInputsCmd.Flags().String("next-token", "", "The token that you can use to return the next set of results.")
	ioteventsCmd.AddCommand(iotevents_listInputsCmd)
}
