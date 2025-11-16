package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lambda_listLayersCmd = &cobra.Command{
	Use:   "list-layers",
	Short: "Lists [Lambda layers](https://docs.aws.amazon.com/lambda/latest/dg/invocation-layers.html) and shows information about the latest version of each.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lambda_listLayersCmd).Standalone()

	lambda_listLayersCmd.Flags().String("compatible-architecture", "", "The compatible [instruction set architecture](https://docs.aws.amazon.com/lambda/latest/dg/foundation-arch.html).")
	lambda_listLayersCmd.Flags().String("compatible-runtime", "", "A runtime identifier.")
	lambda_listLayersCmd.Flags().String("marker", "", "A pagination token returned by a previous call.")
	lambda_listLayersCmd.Flags().String("max-items", "", "The maximum number of layers to return.")
	lambdaCmd.AddCommand(lambda_listLayersCmd)
}
