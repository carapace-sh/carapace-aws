package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appstream_describeFleetsCmd = &cobra.Command{
	Use:   "describe-fleets",
	Short: "Retrieves a list that describes one or more specified fleets, if the fleet names are provided.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appstream_describeFleetsCmd).Standalone()

	appstream_describeFleetsCmd.Flags().String("names", "", "The names of the fleets to describe.")
	appstream_describeFleetsCmd.Flags().String("next-token", "", "The pagination token to use to retrieve the next page of results for this operation.")
	appstreamCmd.AddCommand(appstream_describeFleetsCmd)
}
