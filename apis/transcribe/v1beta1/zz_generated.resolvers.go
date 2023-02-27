/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	v1beta1 "github.com/upbound/provider-aws/apis/iam/v1beta1"
	resource "github.com/upbound/upjet/pkg/resource"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this LanguageModel.
func (mg *LanguageModel) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.InputDataConfig); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.InputDataConfig[i3].DataAccessRoleArn),
			Extract:      resource.ExtractParamPath("arn", true),
			Reference:    mg.Spec.ForProvider.InputDataConfig[i3].DataAccessRoleArnRef,
			Selector:     mg.Spec.ForProvider.InputDataConfig[i3].DataAccessRoleArnSelector,
			To: reference.To{
				List:    &v1beta1.RoleList{},
				Managed: &v1beta1.Role{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.InputDataConfig[i3].DataAccessRoleArn")
		}
		mg.Spec.ForProvider.InputDataConfig[i3].DataAccessRoleArn = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.InputDataConfig[i3].DataAccessRoleArnRef = rsp.ResolvedReference

	}

	return nil
}