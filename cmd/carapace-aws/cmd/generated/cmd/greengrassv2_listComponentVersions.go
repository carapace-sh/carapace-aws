package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var greengrassv2_listComponentVersionsCmd = &cobra.Command{
	Use:   "list-component-versions",
	Short: "Retrieves a paginated list of all versions for a component.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(greengrassv2_listComponentVersionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(greengrassv2_listComponentVersionsCmd).Standalone()

		greengrassv2_listComponentVersionsCmd.Flags().String("arn", "", "The [ARN](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of the component.")
		greengrassv2_listComponentVersionsCmd.Flags().String("max-results", "", "The maximum number of results to be returned per paginated request.")
		greengrassv2_listComponentVersionsCmd.Flags().String("next-token", "", "The token to be used for the next set of paginated results.")
		greengrassv2_listComponentVersionsCmd.MarkFlagRequired("arn")
	})
	greengrassv2Cmd.AddCommand(greengrassv2_listComponentVersionsCmd)
}
