package length

//
// func TestLength_Centimeters(t *testing.T) {
// 	tests := []struct {
// 		name string
// 		s    Length
// 		want uint64
// 	}{
// 		{
// 			name: "Centimeters to Centimeters",
// 			s: Centimeter(10),
// 			want: 10,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			fmt.Sprintf("%v",tt.s)
// 			if got := tt.s.Centimeters(); got != tt.want {
// 				t.Errorf("Centimeters() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func TestLength_Centimeters(t *testing.T) {
// 	tests := []struct {
// 		name string
// 		s    Length
// 		want uint64
// 	}{
// 		// {
// 		// 	name: "Centimeters to Centimeters",
// 		// 	s: Centimeter(10),
// 		// 	want: 10,
// 		// },
// 		{
// 			name: "Millimeters to Centimeters",
// 			s: Millimeters(100),
// 			want: 10,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			fmt.Sprintf("%v",tt.s)
// 			if got := tt.s.Centimeters(); got != tt.want {
// 				t.Errorf("Centimeters() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
//
// func TestLength_Prefix(t *testing.T) {
// 	tests := []struct {
// 		name string
// 		s    Length
// 		want *big.Int
// 	}{
// 		{
// 			name: "Centi",
// 			s: Centimeter(10),
// 			want: big.NewInt(1).Exp(big.NewInt(1), big.NewInt(28), nil),// 1E+28,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			var b *big.Int = tt.s
// 			if  tt.s != tt.want {
// 				t.Errorf("Centimeters() = %v, want %v", b.String(), tt.want.String())
// 			}
// 		})
// 	}
// }