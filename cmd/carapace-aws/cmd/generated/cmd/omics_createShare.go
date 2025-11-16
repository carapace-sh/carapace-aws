package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var omics_createShareCmd = &cobra.Command{
	Use:   "create-share",
	Short: "Creates a cross-account shared resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(omics_createShareCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(omics_createShareCmd).Standalone()

		omics_createShareCmd.Flags().String("principal-subscriber", "", "The principal subscriber is the account being offered shared access to the resource.")
		omics_createShareCmd.Flags().String("resource-arn", "", "The ARN of the resource to be shared.")
		omics_createShareCmd.Flags().String("share-name", "", "A name that the owner defines for the share.")
		omics_createShareCmd.MarkFlagRequired("principal-subscriber")
		omics_createShareCmd.MarkFlagRequired("resource-arn")
	})
	omicsCmd.AddCommand(omics_createShareCmd)
}
