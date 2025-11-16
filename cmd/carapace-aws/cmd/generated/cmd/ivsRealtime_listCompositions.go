package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ivsRealtime_listCompositionsCmd = &cobra.Command{
	Use:   "list-compositions",
	Short: "Gets summary information about all Compositions in your account, in the AWS region where the API request is processed.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ivsRealtime_listCompositionsCmd).Standalone()

	ivsRealtime_listCompositionsCmd.Flags().String("filter-by-encoder-configuration-arn", "", "Filters the Composition list to match the specified EncoderConfiguration attached to at least one of its output.")
	ivsRealtime_listCompositionsCmd.Flags().String("filter-by-stage-arn", "", "Filters the Composition list to match the specified Stage ARN.")
	ivsRealtime_listCompositionsCmd.Flags().String("max-results", "", "Maximum number of results to return.")
	ivsRealtime_listCompositionsCmd.Flags().String("next-token", "", "The first Composition to retrieve.")
	ivsRealtimeCmd.AddCommand(ivsRealtime_listCompositionsCmd)
}
