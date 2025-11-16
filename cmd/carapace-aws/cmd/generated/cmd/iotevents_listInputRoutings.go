package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotevents_listInputRoutingsCmd = &cobra.Command{
	Use:   "list-input-routings",
	Short: "Lists one or more input routings.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotevents_listInputRoutingsCmd).Standalone()

	iotevents_listInputRoutingsCmd.Flags().String("input-identifier", "", "The identifer of the routed input.")
	iotevents_listInputRoutingsCmd.Flags().String("max-results", "", "The maximum number of results to be returned per request.")
	iotevents_listInputRoutingsCmd.Flags().String("next-token", "", "The token that you can use to return the next set of results.")
	iotevents_listInputRoutingsCmd.MarkFlagRequired("input-identifier")
	ioteventsCmd.AddCommand(iotevents_listInputRoutingsCmd)
}
