package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ivs_listStreamsCmd = &cobra.Command{
	Use:   "list-streams",
	Short: "Gets summary information about live streams in your account, in the Amazon Web Services region where the API request is processed.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ivs_listStreamsCmd).Standalone()

	ivs_listStreamsCmd.Flags().String("filter-by", "", "Filters the stream list to match the specified criterion.")
	ivs_listStreamsCmd.Flags().String("max-results", "", "Maximum number of streams to return.")
	ivs_listStreamsCmd.Flags().String("next-token", "", "The first stream to retrieve.")
	ivsCmd.AddCommand(ivs_listStreamsCmd)
}
