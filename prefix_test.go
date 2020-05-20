package uom

import (
	"fmt"
	"testing"
)

func Test_Prefix(t *testing.T) {
	tests := []struct {
		name string
		s    Prefix
		want string
	}{
		{
			name: "Yotta",
			s: Yotta,
			want: "1e+24",
		},
		{
			name: "Zetta",
			s: Zetta,
			want: "1e+21",
		},
		{
			name: "Exa",
			s: Exa,
			want: "1e+18",
		},
		{
			name: "Peta",
			s: Peta,
			want: "1e+15",
		},
		{
			name: "Tera",
			s: Tera,
			want: "1e+12",
		},
		{
			name: "Giga",
			s: Giga,
			want: "1e+09",
		},
		{
			name: "Mega",
			s: Mega,
			want: "1e+06",
		},
		{
			name: "Kilo",
			s: Kilo,
			want: "1000",
		},
		{
			name: "Hecto",
			s: Hecto,
			want: "100",
		},
		{
			name: "Deka",
			s: Deka,
			want: "10",
		},
		{
			name: "Deci",
			s: Deci,
			want: "0.1",
		},
		{
			name: "Centi",
			s: Centi,
			want: "0.01",
		},
		{
			name: "Milli",
			s: Milli,
			want: "0.001",
		},
		{
			name: "Micro",
			s: Micro,
			want: "1e-06",
		},
		{
			name: "Nano",
			s: Nano,
			want: "1e-10",
		},
		{
			name: "Pico",
			s: Pico,
			want: "1e-12",
		},
		{
			name: "Femto",
			s: Femto,
			want: "1e-15",
		},
		{
			name: "Atto",
			s: Atto,
			want: "1e-18",
		},
		{
			name: "Zepto",
			s: Zepto,
			want: "1e-21",
		},
		{
			name: "Yocto",
			s: Yocto,
			want: "1e-24",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if  fmt.Sprintf("%v", tt.s) != tt.want {
				t.Errorf("Centimeters() = %v, want %v", tt.s, tt.want)
			}
		})
	}
}