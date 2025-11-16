package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ds_createHybridAdCmd = &cobra.Command{
	Use:   "create-hybrid-ad",
	Short: "Creates a hybrid directory that connects your self-managed Active Directory (AD) infrastructure and Amazon Web Services.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ds_createHybridAdCmd).Standalone()

	ds_createHybridAdCmd.Flags().String("assessment-id", "", "The unique identifier of the successful directory assessment that validates your self-managed AD environment.")
	ds_createHybridAdCmd.Flags().String("secret-arn", "", "The Amazon Resource Name (ARN) of the Amazon Web Services Secrets Manager secret that contains the credentials for the service account used to join hybrid domain controllers to your self-managed AD domain.")
	ds_createHybridAdCmd.Flags().String("tags", "", "The tags to be assigned to the directory.")
	ds_createHybridAdCmd.MarkFlagRequired("assessment-id")
	ds_createHybridAdCmd.MarkFlagRequired("secret-arn")
	dsCmd.AddCommand(ds_createHybridAdCmd)
}
