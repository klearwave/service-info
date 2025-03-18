package models

import (
	"testing"

	"github.com/klearwave/service-info/pkg/models"
	"github.com/klearwave/service-info/pkg/utils/pointers"
)

const (
	testContainerImageRegistry         = "ghcr.io"
	testContainerImageRepo             = "klearwave"
	testContainerImageMinimal          = "postgres"
	testContainerImageTag              = "v0.1.2"
	testContainerImageSHA256SumValid   = "2d4b92db6941294f731cfe7aeca336eb8dba279171c0e6ceda32b9f018f8429d"
	testContainerImageSHA256SumInvalid = "2d4b92db6941294f731cfe7aeca336eb8dba279171c0e6ceda32b9f018f8429e"
	testContainerImageCommitHash       = "631af50a8bbc4b5e69dab77d51a3a1733550fe8d"

	testContainerImageMinimalWithTag      = testContainerImageMinimal + ":" + testContainerImageTag
	testContainerImageWithRegistry        = testContainerImageRegistry + "/" + testContainerImageRepo + "/" + testContainerImageMinimal
	testContainerImageWithRegistryWithTag = testContainerImageWithRegistry + ":" + testContainerImageTag
	testContainerImageWithRegistryFull    = testContainerImageWithRegistryWithTag + "@sha256:" + testContainerImageSHA256SumValid
)

func TestContainerImage_Parse(t *testing.T) {
	t.Parallel()

	type fields struct {
		Model              models.Model
		ContainerImageBase ContainerImageBase
		ImageRegistry      *string
		ImageName          *string
		ImageTag           *string
	}

	tests := []struct {
		name    string
		fields  fields
		wantErr bool
		want    *ContainerImage
	}{
		{
			name: "fail: missing container image",
			fields: fields{
				ContainerImageBase: ContainerImageBase{},
			},
			wantErr: true,
		},
		{
			name: "fail: invalid container image",
			fields: fields{
				ContainerImageBase: ContainerImageBase{
					Image: pointers.FromString("thisisfake@latest"),
				},
			},
			wantErr: true,
		},
		{
			name: "fail: mismatch value",
			fields: fields{
				ContainerImageBase: ContainerImageBase{
					Image:     pointers.FromString(testContainerImageWithRegistryFull),
					SHA256Sum: pointers.FromString(testContainerImageSHA256SumInvalid),
				},
			},
			wantErr: true,
		},
		{
			name: "fail: missing required value",
			fields: fields{
				ContainerImageBase: ContainerImageBase{
					Image: pointers.FromString(testContainerImageWithRegistryWithTag),
				},
			},
			wantErr: true,
		},
		{
			name: "success: ensure minimal image is parsed correctly",
			fields: fields{
				ContainerImageBase: ContainerImageBase{
					Image:      pointers.FromString(testContainerImageMinimal),
					SHA256Sum:  pointers.FromString(testContainerImageSHA256SumValid),
					CommitHash: pointers.FromString(testContainerImageCommitHash),
				},
			},
			want: &ContainerImage{
				ContainerImageBase: ContainerImageBase{
					Image:      pointers.FromString(testContainerImageMinimal),
					SHA256Sum:  pointers.FromString(testContainerImageSHA256SumValid),
					CommitHash: pointers.FromString(testContainerImageCommitHash),
				},
				ImageRegistry: pointers.FromString(""),
				ImageName:     pointers.FromString(testContainerImageMinimal),
				ImageTag:      pointers.FromString(""),
			},
		},
		{
			name: "success: ensure minimal image with tag is parsed correctly",
			fields: fields{
				ContainerImageBase: ContainerImageBase{
					Image:      pointers.FromString(testContainerImageMinimalWithTag),
					SHA256Sum:  pointers.FromString(testContainerImageSHA256SumValid),
					CommitHash: pointers.FromString(testContainerImageCommitHash),
				},
			},
			want: &ContainerImage{
				ContainerImageBase: ContainerImageBase{
					Image:      pointers.FromString(testContainerImageMinimalWithTag),
					SHA256Sum:  pointers.FromString(testContainerImageSHA256SumValid),
					CommitHash: pointers.FromString(testContainerImageCommitHash),
				},
				ImageRegistry: pointers.FromString(""),
				ImageName:     pointers.FromString(testContainerImageMinimal),
				ImageTag:      pointers.FromString(testContainerImageTag),
			},
		},
		{
			name: "success: ensure image with registry without tag is parsed correctly",
			fields: fields{
				ContainerImageBase: ContainerImageBase{
					Image:      pointers.FromString(testContainerImageWithRegistry),
					SHA256Sum:  pointers.FromString(testContainerImageSHA256SumValid),
					CommitHash: pointers.FromString(testContainerImageCommitHash),
				},
			},
			want: &ContainerImage{
				ContainerImageBase: ContainerImageBase{
					Image:      pointers.FromString(testContainerImageWithRegistry),
					SHA256Sum:  pointers.FromString(testContainerImageSHA256SumValid),
					CommitHash: pointers.FromString(testContainerImageCommitHash),
				},
				ImageRegistry: pointers.FromString(testContainerImageRegistry),
				ImageName:     pointers.FromString(testContainerImageRepo + "/" + testContainerImageMinimal),
				ImageTag:      pointers.FromString(""),
			},
		},
		{
			name: "success: ensure image with registry with tag is parsed correctly",
			fields: fields{
				ContainerImageBase: ContainerImageBase{
					Image:      pointers.FromString(testContainerImageWithRegistryWithTag),
					SHA256Sum:  pointers.FromString(testContainerImageSHA256SumValid),
					CommitHash: pointers.FromString(testContainerImageCommitHash),
				},
			},
			want: &ContainerImage{
				ContainerImageBase: ContainerImageBase{
					Image:      pointers.FromString(testContainerImageWithRegistryWithTag),
					SHA256Sum:  pointers.FromString(testContainerImageSHA256SumValid),
					CommitHash: pointers.FromString(testContainerImageCommitHash),
				},
				ImageRegistry: pointers.FromString(testContainerImageRegistry),
				ImageName:     pointers.FromString(testContainerImageRepo + "/" + testContainerImageMinimal),
				ImageTag:      pointers.FromString(testContainerImageTag),
			},
		},
		{
			name: "success: ensure full image is parsed correctly",
			fields: fields{
				ContainerImageBase: ContainerImageBase{
					Image:      pointers.FromString(testContainerImageWithRegistryFull),
					CommitHash: pointers.FromString(testContainerImageCommitHash),
				},
			},
			want: &ContainerImage{
				ContainerImageBase: ContainerImageBase{
					Image:      pointers.FromString(testContainerImageWithRegistryFull),
					SHA256Sum:  pointers.FromString(testContainerImageSHA256SumValid),
					CommitHash: pointers.FromString(testContainerImageCommitHash),
				},
				ImageRegistry: pointers.FromString(testContainerImageRegistry),
				ImageName:     pointers.FromString(testContainerImageRepo + "/" + testContainerImageMinimal),
				ImageTag:      pointers.FromString(testContainerImageTag),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			containerImage := &ContainerImage{
				Model:              tt.fields.Model,
				ContainerImageBase: tt.fields.ContainerImageBase,
				ImageRegistry:      tt.fields.ImageRegistry,
				ImageName:          tt.fields.ImageName,
				ImageTag:           tt.fields.ImageTag,
			}
			if err := containerImage.Parse(); (err != nil) != tt.wantErr {
				t.Errorf("ContainerImage.Parse() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.want != nil {
				for want, got := range map[*string]*string{
					tt.want.SHA256Sum:     containerImage.SHA256Sum,
					tt.want.CommitHash:    containerImage.CommitHash,
					tt.want.Image:         containerImage.Image,
					tt.want.ImageRegistry: containerImage.ImageRegistry,
					tt.want.ImageName:     containerImage.ImageName,
					tt.want.ImageTag:      containerImage.ImageTag,
				} {
					wantValue := ""
					gotValue := ""

					if want != nil {
						wantValue = *want
					}

					if got != nil {
						gotValue = *got
					}

					if !pointers.EqualString(want, got) {
						t.Errorf("ContainerImage.Parse(); want value = [%s], got value [%s]", wantValue, gotValue)
					}
				}
			}
		})
	}
}
