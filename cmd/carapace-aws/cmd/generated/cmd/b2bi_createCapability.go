package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var b2bi_createCapabilityCmd = &cobra.Command{
	Use:   "create-capability",
	Short: "Instantiates a capability based on the specified parameters.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(b2bi_createCapabilityCmd).Standalone()

	b2bi_createCapabilityCmd.Flags().String("client-token", "", "Reserved for future use.")
	b2bi_createCapabilityCmd.Flags().String("configuration", "", "Specifies a structure that contains the details for a capability.")
	b2bi_createCapabilityCmd.Flags().String("instructions-documents", "", "Specifies one or more locations in Amazon S3, each specifying an EDI document that can be used with this capability.")
	b2bi_createCapabilityCmd.Flags().String("name", "", "Specifies the name of the capability, used to identify it.")
	b2bi_createCapabilityCmd.Flags().String("tags", "", "Specifies the key-value pairs assigned to ARNs that you can use to group and search for resources by type.")
	b2bi_createCapabilityCmd.Flags().String("type", "", "Specifies the type of the capability.")
	b2bi_createCapabilityCmd.MarkFlagRequired("configuration")
	b2bi_createCapabilityCmd.MarkFlagRequired("name")
	b2bi_createCapabilityCmd.MarkFlagRequired("type")
	b2biCmd.AddCommand(b2bi_createCapabilityCmd)
}
