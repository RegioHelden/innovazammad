package innovaphone

import (
	"testing"

	"github.com/regiohelden/innovazammad/config"
)

func TestNo_Normalize(t *testing.T) {
	type fields struct {
		Type string
		Cn   string
		Dn   string
		E164 string
		H323 string
	}
	tests := []struct {
		name   string
		fields fields
		config config.Config
		wantN  string
	}{
		{
			"internal without prefix",
			fields{Type: "this", E164: "123"},
			config.Config{},
			"123",
		},
		{
			"internal to prefixed",
			fields{Type: "this", E164: "123"},
			config.Config{
				Zammad: config.Zammad{NumberPrefix: "49456", CountryCode: 49},
			},
			"49456123",
		},
		{
			"external with zero prefix",
			fields{Type: "peer", E164: "00123456789"},
			config.Config{
				Zammad: config.Zammad{TrunkPrefix: "0", NumberPrefix: "49456", CountryCode: 49},
			},
			"49123456789",
		},
		{
			"external pre-normalized",
			fields{Type: "peer", E164: "049123456789"},
			config.Config{
				Zammad: config.Zammad{TrunkPrefix: "0", NumberPrefix: "49456", CountryCode: 49},
			},
			"49123456789",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config.Global = tt.config
			no := &No{
				Type: tt.fields.Type,
				Cn:   tt.fields.Cn,
				Dn:   tt.fields.Dn,
				E164: tt.fields.E164,
				H323: tt.fields.H323,
			}
			if gotN := no.Normalize(); gotN != tt.wantN {
				t.Errorf("No.Normalize() = %v, want %v", gotN, tt.wantN)
			}
		})
	}
}
