package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi-tailscale/sdk/go/tailscale"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := tailscale.NewDnsNameservers(ctx, "NameServers", &tailscale.DnsNameserversArgs{
			Nameservers: pulumi.StringArray{
                pulumi.String("1.1.1.1"),
				pulumi.String("1.0.0.1"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
