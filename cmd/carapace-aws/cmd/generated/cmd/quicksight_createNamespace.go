package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_createNamespaceCmd = &cobra.Command{
	Use:   "create-namespace",
	Short: "(Enterprise edition only) Creates a new namespace for you to use with Amazon Quick Sight.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_createNamespaceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(quicksight_createNamespaceCmd).Standalone()

		quicksight_createNamespaceCmd.Flags().String("aws-account-id", "", "The ID for the Amazon Web Services account that you want to create the Quick Sight namespace in.")
		quicksight_createNamespaceCmd.Flags().String("identity-store", "", "Specifies the type of your user identity directory.")
		quicksight_createNamespaceCmd.Flags().String("namespace", "", "The name that you want to use to describe the new namespace.")
		quicksight_createNamespaceCmd.Flags().String("tags", "", "The tags that you want to associate with the namespace that you're creating.")
		quicksight_createNamespaceCmd.MarkFlagRequired("aws-account-id")
		quicksight_createNamespaceCmd.MarkFlagRequired("identity-store")
		quicksight_createNamespaceCmd.MarkFlagRequired("namespace")
	})
	quicksightCmd.AddCommand(quicksight_createNamespaceCmd)
}
