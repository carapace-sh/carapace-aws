package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elastictranscoder_cancelJobCmd = &cobra.Command{
	Use:   "cancel-job",
	Short: "The CancelJob operation cancels an unfinished job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elastictranscoder_cancelJobCmd).Standalone()

	elastictranscoder_cancelJobCmd.Flags().String("id", "", "The identifier of the job that you want to cancel.")
	elastictranscoder_cancelJobCmd.MarkFlagRequired("id")
	elastictranscoderCmd.AddCommand(elastictranscoder_cancelJobCmd)
}
