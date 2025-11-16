package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediaconvert_listQueuesCmd = &cobra.Command{
	Use:   "list-queues",
	Short: "Retrieve a JSON array of up to twenty of your queues.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediaconvert_listQueuesCmd).Standalone()

	mediaconvert_listQueuesCmd.Flags().String("list-by", "", "Optional.")
	mediaconvert_listQueuesCmd.Flags().String("max-results", "", "Optional.")
	mediaconvert_listQueuesCmd.Flags().String("next-token", "", "Use this string, provided with the response to a previous request, to request the next batch of queues.")
	mediaconvert_listQueuesCmd.Flags().String("order", "", "Optional.")
	mediaconvertCmd.AddCommand(mediaconvert_listQueuesCmd)
}
