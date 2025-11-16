package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kms_createCustomKeyStoreCmd = &cobra.Command{
	Use:   "create-custom-key-store",
	Short: "Creates a [custom key store](https://docs.aws.amazon.com/kms/latest/developerguide/key-store-overview.html) backed by a key store that you own and manage.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kms_createCustomKeyStoreCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kms_createCustomKeyStoreCmd).Standalone()

		kms_createCustomKeyStoreCmd.Flags().String("cloud-hsm-cluster-id", "", "Identifies the CloudHSM cluster for an CloudHSM key store.")
		kms_createCustomKeyStoreCmd.Flags().String("custom-key-store-name", "", "Specifies a friendly name for the custom key store.")
		kms_createCustomKeyStoreCmd.Flags().String("custom-key-store-type", "", "Specifies the type of custom key store.")
		kms_createCustomKeyStoreCmd.Flags().String("key-store-password", "", "Specifies the `kmsuser` password for an CloudHSM key store.")
		kms_createCustomKeyStoreCmd.Flags().String("trust-anchor-certificate", "", "Specifies the certificate for an CloudHSM key store.")
		kms_createCustomKeyStoreCmd.Flags().String("xks-proxy-authentication-credential", "", "Specifies an authentication credential for the external key store proxy (XKS proxy).")
		kms_createCustomKeyStoreCmd.Flags().String("xks-proxy-connectivity", "", "Indicates how KMS communicates with the external key store proxy.")
		kms_createCustomKeyStoreCmd.Flags().String("xks-proxy-uri-endpoint", "", "Specifies the endpoint that KMS uses to send requests to the external key store proxy (XKS proxy).")
		kms_createCustomKeyStoreCmd.Flags().String("xks-proxy-uri-path", "", "Specifies the base path to the proxy APIs for this external key store.")
		kms_createCustomKeyStoreCmd.Flags().String("xks-proxy-vpc-endpoint-service-name", "", "Specifies the name of the Amazon VPC endpoint service for interface endpoints that is used to communicate with your external key store proxy (XKS proxy).")
		kms_createCustomKeyStoreCmd.Flags().String("xks-proxy-vpc-endpoint-service-owner", "", "Specifies the Amazon Web Services account ID that owns the Amazon VPC service endpoint for the interface that is used to communicate with your external key store proxy (XKS proxy).")
		kms_createCustomKeyStoreCmd.MarkFlagRequired("custom-key-store-name")
	})
	kmsCmd.AddCommand(kms_createCustomKeyStoreCmd)
}
