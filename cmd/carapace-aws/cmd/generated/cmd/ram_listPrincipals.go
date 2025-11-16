package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ram_listPrincipalsCmd = &cobra.Command{
	Use:   "list-principals",
	Short: "Lists the principals that you are sharing resources with or that are sharing resources with you.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ram_listPrincipalsCmd).Standalone()

	ram_listPrincipalsCmd.Flags().String("max-results", "", "Specifies the total number of results that you want included on each page of the response.")
	ram_listPrincipalsCmd.Flags().String("next-token", "", "Specifies that you want to receive the next page of results.")
	ram_listPrincipalsCmd.Flags().String("principals", "", "Specifies that you want to list information for only the listed principals.")
	ram_listPrincipalsCmd.Flags().String("resource-arn", "", "Specifies that you want to list principal information for the resource share with the specified [Amazon Resource Name (ARN)](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html).")
	ram_listPrincipalsCmd.Flags().String("resource-owner", "", "Specifies that you want to list information for only resource shares that match the following:")
	ram_listPrincipalsCmd.Flags().String("resource-share-arns", "", "Specifies that you want to list information for only principals associated with the resource shares specified by a list the [Amazon Resource Names (ARNs)](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html).")
	ram_listPrincipalsCmd.Flags().String("resource-type", "", "Specifies that you want to list information for only principals associated with resource shares that include the specified resource type.")
	ram_listPrincipalsCmd.MarkFlagRequired("resource-owner")
	ramCmd.AddCommand(ram_listPrincipalsCmd)
}
