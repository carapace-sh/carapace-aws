package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockRuntime_listAsyncInvokesCmd = &cobra.Command{
	Use:   "list-async-invokes",
	Short: "Lists asynchronous invocations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockRuntime_listAsyncInvokesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrockRuntime_listAsyncInvokesCmd).Standalone()

		bedrockRuntime_listAsyncInvokesCmd.Flags().String("max-results", "", "The maximum number of invocations to return in one page of results.")
		bedrockRuntime_listAsyncInvokesCmd.Flags().String("next-token", "", "Specify the pagination token from a previous request to retrieve the next page of results.")
		bedrockRuntime_listAsyncInvokesCmd.Flags().String("sort-by", "", "How to sort the response.")
		bedrockRuntime_listAsyncInvokesCmd.Flags().String("sort-order", "", "The sorting order for the response.")
		bedrockRuntime_listAsyncInvokesCmd.Flags().String("status-equals", "", "Filter invocations by status.")
		bedrockRuntime_listAsyncInvokesCmd.Flags().String("submit-time-after", "", "Include invocations submitted after this time.")
		bedrockRuntime_listAsyncInvokesCmd.Flags().String("submit-time-before", "", "Include invocations submitted before this time.")
	})
	bedrockRuntimeCmd.AddCommand(bedrockRuntime_listAsyncInvokesCmd)
}
