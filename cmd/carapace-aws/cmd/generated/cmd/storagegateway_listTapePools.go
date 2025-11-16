package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var storagegateway_listTapePoolsCmd = &cobra.Command{
	Use:   "list-tape-pools",
	Short: "Lists custom tape pools.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(storagegateway_listTapePoolsCmd).Standalone()

	storagegateway_listTapePoolsCmd.Flags().String("limit", "", "An optional number limit for the tape pools in the list returned by this call.")
	storagegateway_listTapePoolsCmd.Flags().String("marker", "", "A string that indicates the position at which to begin the returned list of tape pools.")
	storagegateway_listTapePoolsCmd.Flags().String("pool-arns", "", "The Amazon Resource Name (ARN) of each of the custom tape pools you want to list.")
	storagegatewayCmd.AddCommand(storagegateway_listTapePoolsCmd)
}
