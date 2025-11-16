package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_listAliasesCmd = &cobra.Command{
	Use:   "list-aliases",
	Short: "Lists the aliases of a specified image or image version.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_listAliasesCmd).Standalone()

	sagemaker_listAliasesCmd.Flags().String("alias", "", "The alias of the image version.")
	sagemaker_listAliasesCmd.Flags().String("image-name", "", "The name of the image.")
	sagemaker_listAliasesCmd.Flags().String("max-results", "", "The maximum number of aliases to return.")
	sagemaker_listAliasesCmd.Flags().String("next-token", "", "If the previous call to `ListAliases` didn't return the full set of aliases, the call returns a token for retrieving the next set of aliases.")
	sagemaker_listAliasesCmd.Flags().String("version", "", "The version of the image.")
	sagemaker_listAliasesCmd.MarkFlagRequired("image-name")
	sagemakerCmd.AddCommand(sagemaker_listAliasesCmd)
}
