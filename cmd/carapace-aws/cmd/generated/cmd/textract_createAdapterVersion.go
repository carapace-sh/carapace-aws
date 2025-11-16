package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var textract_createAdapterVersionCmd = &cobra.Command{
	Use:   "create-adapter-version",
	Short: "Creates a new version of an adapter.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(textract_createAdapterVersionCmd).Standalone()

	textract_createAdapterVersionCmd.Flags().String("adapter-id", "", "A string containing a unique ID for the adapter that will receive a new version.")
	textract_createAdapterVersionCmd.Flags().String("client-request-token", "", "Idempotent token is used to recognize the request.")
	textract_createAdapterVersionCmd.Flags().String("dataset-config", "", "Specifies a dataset used to train a new adapter version.")
	textract_createAdapterVersionCmd.Flags().String("kmskey-id", "", "The identifier for your AWS Key Management Service key (AWS KMS key).")
	textract_createAdapterVersionCmd.Flags().String("output-config", "", "")
	textract_createAdapterVersionCmd.Flags().String("tags", "", "A set of tags (key-value pairs) that you want to attach to the adapter version.")
	textract_createAdapterVersionCmd.MarkFlagRequired("adapter-id")
	textract_createAdapterVersionCmd.MarkFlagRequired("dataset-config")
	textract_createAdapterVersionCmd.MarkFlagRequired("output-config")
	textractCmd.AddCommand(textract_createAdapterVersionCmd)
}
