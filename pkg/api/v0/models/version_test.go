package models

import (
	"reflect"
	"testing"

	"gorm.io/gorm"

	"github.com/klearwave/service-info/pkg/models"
	"github.com/klearwave/service-info/pkg/utils/pointers"
)

func TestVersion_BeforeCreate(t *testing.T) {
	t.Parallel()

	type fields struct {
		Model           models.Model
		VersionBase     VersionBase
		Stable          *bool
		XVersion        *int
		YVersion        *int
		ZVersion        *int
		BuildVersion    *string
		ContainerImages []*ContainerImage
	}

	type args struct {
		tx *gorm.DB
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		want    *Version
	}{
		{
			name: "fail: missing version id",
			fields: fields{
				VersionBase: VersionBase{},
			},
			wantErr: true,
		},
		{
			name: "fail: invalid version id",
			fields: fields{
				VersionBase: VersionBase{
					Id: pointers.FromString("v1.2"),
				},
			},
			wantErr: true,
		},
		{
			name: "success: v is prepended",
			fields: fields{
				VersionBase: VersionBase{
					Id: pointers.FromString("1.2.3"),
				},
			},
			want: &Version{
				VersionBase: VersionBase{
					Id: pointers.FromString("v1.2.3"),
				},
				XVersion: pointers.Int(1),
				YVersion: pointers.Int(2),
				ZVersion: pointers.Int(3),
				Stable:   pointers.Bool(true),
			},
		},
		{
			name: "success: x/y/z versions are set",
			fields: fields{
				VersionBase: VersionBase{
					Id: pointers.FromString("v1.2.3"),
				},
			},
			want: &Version{
				VersionBase: VersionBase{
					Id: pointers.FromString("v1.2.3"),
				},
				XVersion: pointers.Int(1),
				YVersion: pointers.Int(2),
				ZVersion: pointers.Int(3),
				Stable:   pointers.Bool(true),
			},
		},
		{
			name: "success: build version is set",
			fields: fields{
				VersionBase: VersionBase{
					Id: pointers.FromString("v1.2.3-prerelease.1"),
				},
			},
			want: &Version{
				VersionBase: VersionBase{
					Id: pointers.FromString("v1.2.3-prerelease.1"),
				},
				XVersion:     pointers.Int(1),
				YVersion:     pointers.Int(2),
				ZVersion:     pointers.Int(3),
				BuildVersion: pointers.FromString("prerelease.1"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			version := &Version{
				Model:           tt.fields.Model,
				VersionBase:     tt.fields.VersionBase,
				Stable:          tt.fields.Stable,
				XVersion:        tt.fields.XVersion,
				YVersion:        tt.fields.YVersion,
				ZVersion:        tt.fields.ZVersion,
				BuildVersion:    tt.fields.BuildVersion,
				ContainerImages: tt.fields.ContainerImages,
			}
			if err := version.BeforeCreate(tt.args.tx); (err != nil) != tt.wantErr {
				t.Errorf("Version.BeforeCreate() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.want != nil && !reflect.DeepEqual(version, tt.want) {
				t.Errorf("Version got = %+v, want %+v", version, tt.want)
			}
		})
	}
}
