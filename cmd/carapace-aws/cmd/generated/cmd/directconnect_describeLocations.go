package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var directconnect_describeLocationsCmd = &cobra.Command{
	Use:   "describe-locations",
	Short: "Lists the Direct Connect locations in the current Amazon Web Services Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(directconnect_describeLocationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(directconnect_describeLocationsCmd).Standalone()

	})
	directconnectCmd.AddCommand(directconnect_describeLocationsCmd)
}
