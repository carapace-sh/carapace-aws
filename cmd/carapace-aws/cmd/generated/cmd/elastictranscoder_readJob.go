package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elastictranscoder_readJobCmd = &cobra.Command{
	Use:   "read-job",
	Short: "The ReadJob operation returns detailed information about a job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elastictranscoder_readJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(elastictranscoder_readJobCmd).Standalone()

		elastictranscoder_readJobCmd.Flags().String("id", "", "The identifier of the job for which you want to get detailed information.")
		elastictranscoder_readJobCmd.MarkFlagRequired("id")
	})
	elastictranscoderCmd.AddCommand(elastictranscoder_readJobCmd)
}
